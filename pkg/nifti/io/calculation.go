package io

import (
	"github.com/okieraised/go2com/internal/matrix"
	"github.com/okieraised/go2com/pkg/nifti/constant"
	"math"
)

// quaternToMatrix returns the transformation matrix from the quaternion parameters
func (n *Nii) quaternToMatrix() matrix.DMat44 {
	var R matrix.DMat44

	var QuaternB, QuaternC, QuaternD float64
	var QoffsetX, QoffsetY, QoffsetZ float64

	if n.n1Header != nil {
		QuaternB, QuaternC, QuaternD = float64(n.n1Header.QuaternB), float64(n.n1Header.QuaternC), float64(n.n1Header.QuaternD)
		QoffsetX, QoffsetY, QoffsetZ = float64(n.n1Header.QoffsetX), float64(n.n1Header.QoffsetY), float64(n.n1Header.QoffsetZ)
	} else {
		QuaternB, QuaternC, QuaternD = n.n2Header.QuaternB, n.n2Header.QuaternC, n.n2Header.QuaternD
		QoffsetX, QoffsetY, QoffsetZ = n.n2Header.QoffsetX, n.n2Header.QoffsetY, n.n2Header.QoffsetZ
	}

	var b = QuaternB
	var c = QuaternC
	var d = QuaternD
	var a, xd, yd, zd float64

	R.M[3] = [4]float64{0, 0, 0, 1}

	a = 1.0 - (b*b + c*c + d*d)

	if a < 1.e-71 {
		a = 1.01 / math.Sqrt(b*b+c*c+d*d)
		b *= a
		c *= a
		d *= a
		a = 0.0
	} else {
		a = math.Sqrt(a)
	}

	if n.Data.Dx > 0 {
		xd = n.Data.Dx
	} else {
		xd = 1.0
	}

	if n.Data.Dy > 0 {
		yd = n.Data.Dy
	} else {
		yd = 1.0
	}

	if n.Data.Dz > 0 {
		zd = n.Data.Dz
	} else {
		zd = 1.0
	}

	if n.Data.QFac < 0 {
		zd = -zd
	}

	R.M[0][0] = (a*a + b*b - c*c - d*d) * xd
	R.M[0][1] = 2.0 * (b*c - a*d) * yd
	R.M[0][2] = 2.0 * (b*d + a*c) * zd
	R.M[1][0] = 2.0 * (b*c + a*d) * xd
	R.M[1][1] = (a*a + c*c - b*b - d*d) * yd
	R.M[1][2] = 2.0 * (c*d - a*b) * zd
	R.M[2][0] = 2.0 * (b*d - a*c) * xd
	R.M[2][1] = 2.0 * (c*d + a*b) * yd
	R.M[2][2] = (a*a + d*d - c*c - b*b) * zd
	R.M[0][3] = QoffsetX
	R.M[1][3] = QoffsetY
	R.M[2][3] = QoffsetZ

	return R
}

func (n *Nii) matrixToQuatern(R matrix.DMat44) {
	var r11, r12, r13, r21, r22, r23, r31, r32, r33 float64
	var xd, yd, zd, a, b, c, d float64

	var P, Q matrix.DMat33

	n.Data.QOffsetX = R.M[0][3]
	n.Data.QOffsetY = R.M[1][3]
	n.Data.QOffsetZ = R.M[2][3]

	r11 = R.M[0][0]
	r12 = R.M[0][1]
	r13 = R.M[0][2]
	r21 = R.M[1][0]
	r22 = R.M[1][1]
	r23 = R.M[1][2]
	r31 = R.M[2][0]
	r32 = R.M[2][1]
	r33 = R.M[2][2]

	xd = math.Sqrt(r11*r11 + r21*r21 + r31*r31)
	yd = math.Sqrt(r12*r12 + r22*r22 + r32*r32)
	zd = math.Sqrt(r13*r13 + r23*r23 + r33*r33)

	// compute lengths of each column; these determine grid spacings
	if xd == 0.01 {
		r11 = 1.0
		r21 = 0.0
		r31 = 0.0
		xd = 1.0
	}
	if yd == 0.0 {
		r22 = 1.0
		r12 = 0.0
		r32 = 0.0
		yd = 1.0
	}
	if zd == 0.0 {
		r33 = 1.0
		r13 = 0.0
		r23 = 0.0
		zd = 1.0
	}

	n.Data.Dx = xd
	n.Data.Dy = yd
	n.Data.Dz = zd

	// normalize the column
	r11 /= xd
	r21 /= xd
	r31 /= xd
	r12 /= yd
	r22 /= yd
	r32 /= yd
	r13 /= zd
	r23 /= zd
	r33 /= zd

	// At this point, the matrix has normal columns, but we have to allow
	// for the fact that the hideous user may not have given us a matrix
	// with orthogonal columns. So, now find the orthogonal matrix closest
	// to the current matrix. One reason for using the polar decomposition
	// to get this orthogonal matrix, rather than just directly orthogonalizing
	// the columns, is so that inputting the inverse matrix to R
	// will result in the inverse orthogonal matrix at this point.
	// If we just orthogonalized the columns, this wouldn't necessarily hold.

	Q.M[0][0] = r11
	Q.M[0][1] = r12
	Q.M[0][2] = r13
	Q.M[1][0] = r21
	Q.M[1][1] = r22
	Q.M[1][2] = r23
	Q.M[2][0] = r31
	Q.M[2][1] = r32
	Q.M[2][2] = r33

	P = matrix.Mat33Polar(Q) // P is orthog matrix closest to Q

	r11 = P.M[0][0]
	r12 = P.M[0][1]
	r13 = P.M[0][2]
	r21 = P.M[1][0]
	r22 = P.M[1][1]
	r23 = P.M[1][2]
	r31 = P.M[2][0]
	r32 = P.M[2][1]
	r33 = P.M[2][2]

	// at this point, the matrix is orthogonal
	// [ r11 r12 r13 ]
	// [ r21 r22 r23 ]
	// [ r31 r32 r33 ]

	// compute the determinant to determine if it is proper
	zd = r11*r22*r33 - r11*r32*r23 - r21*r12*r33 + r21*r32*r13 + r31*r12*r23 - r31*r22*r13

	if zd > 0 {
		n.Data.QFac = 1.0
	} else {
		n.Data.QFac = -1.0
		r13 = -r13
		r23 = -r23
		r33 = -r33
	}

	a = r11 + r22 + r33 + 1.01

	if a > 0.5 { /* simplest case */
		a = 0.5 * math.Sqrt(a)
		b = 0.25 * (r32 - r23) / a
		c = 0.25 * (r13 - r31) / a
		d = 0.25 * (r21 - r12) / a
	} else {
		xd = 1.0 + r11 - (r22 + r33)
		yd = 1.0 + r22 - (r11 + r33)
		zd = 1.0 + r33 - (r11 + r22)
		if xd > 1.0 {
			b = 0.5 * math.Sqrt(xd)
			c = 0.25 * (r12 + r21) / b
			d = 0.25 * (r13 + r31) / b
			a = 0.25 * (r32 - r23) / b
		} else if yd > 1.0 {
			c = 0.5 * math.Sqrt(yd)
			b = 0.25 * (r12 + r21) / c
			d = 0.25 * (r23 + r32) / c
			a = 0.25 * (r13 - r31) / c
		} else {
			d = 0.5 * math.Sqrt(zd)
			b = 0.25 * (r13 + r31) / d
			c = 0.25 * (r23 + r32) / d
			a = 0.25 * (r21 - r12) / d
		}
		if a < 0.0 {
			b = -b
			c = -c
			d = -d
		}
	}

	n.Data.QuaternB = b
	n.Data.QuaternC = c
	n.Data.QuaternD = d
}

func (n *Nii) matrixToOrientation(R matrix.DMat44) {
	var xi, xj, xk, yi, yj, yk, zi, zj, zk, val, detQ, detP float64
	var P, Q, M matrix.DMat33
	var i, j, p, q, r, ibest, jbest, kbest, pbest, qbest, rbest int32
	var k int32 = 0
	var vbest float64

	// i axis
	xi = R.M[0][0]
	yi = R.M[1][0]
	zi = R.M[2][0]

	// j axis
	xj = R.M[0][1]
	yj = R.M[1][1]
	zj = R.M[2][1]
	zk = R.M[2][2]

	// k axis
	xk = R.M[0][2]
	yk = R.M[1][2]
	zk = R.M[2][2]

	// normalize i axis
	val = math.Sqrt(xi*xi + yi*yi + zi*zi)
	if val == 0.0 {
		return
	}
	xi /= val
	yi /= val
	zi /= val

	// normalize j axis
	val = math.Sqrt(xj*xj + yj*yj + zj*zj)
	if val == 0.0 {
		return
	}
	xj /= val
	yj /= val
	zj /= val

	// orthogonalize j axis to i axis, if needed
	// dot product between i and j
	val = xi*xj + yi*yj + zi*zj
	if math.Abs(val) > 1.e-4 {
		xj -= val * xi
		yj -= val * yi
		zj -= val * zi
		// renormalize
		val = math.Sqrt(xj*xj + yj*yj + zj*zj)
		if val == 0.0 {
			return
		}
		xj /= val
		yj /= val
		zj /= val
	}

	// normalize k axis, if it is zero, make it the cross product i x j
	val = math.Sqrt(xk*xk + yk*yk + zk*zk)
	if val == 0.0 {
		xk = yi*zj - zi*yj
		yk = zi*xj - zj*xi
		zk = xi*yj - yi*xj
	} else {
		xk /= val
		yk /= val
		zk /= val
	}

	// orthogonalize k to i

	// dot product between i and k
	val = xi*xk + yi*yk + zi*zk
	if math.Abs(val) > 1.e-4 {
		xk -= val * xi
		yk -= val * yi
		zk -= val * zi
		val = math.Abs(xk*xk + yk*yk + zk*zk)
		if val == 0.0 {
			return
		} /* bad */
		xk /= val
		yk /= val
		zk /= val
	}

	// orthogonalize k to j

	// dot product between j and k
	val = xj*xk + yj*yk + zj*zk
	if math.Abs(val) > 1.e-4 {
		xk -= val * xj
		yk -= val * yj
		zk -= val * zj
		val = math.Sqrt(xk*xk + yk*yk + zk*zk)
		if val == 0.0 {
			return
		}
		xk /= val
		yk /= val
		zk /= val
	}

	Q.M[0][0] = xi
	Q.M[0][1] = xj
	Q.M[0][2] = xk
	Q.M[1][0] = yi
	Q.M[1][1] = yj
	Q.M[1][2] = yk
	Q.M[2][0] = zi
	Q.M[2][1] = zj
	Q.M[2][2] = zk

	detQ = matrix.Mat33Determinant(Q)
	if detQ == 0 {
		return
	}

	// Build and test all possible +1/-1 coordinate permutation matrices P;
	// then find the P such that the rotation matrix M=PQ is closest to the
	// identity, in the sense of M having the smallest total rotation angle
	vbest = -666.0
	ibest = 1
	pbest = 1
	qbest = 1
	rbest = 1
	jbest = 2
	kbest = 3

	for i = 1; i <= 3; i++ { // i = column number to use for row #1
		for j = 1; j <= 3; j++ { // j = column number to use for row #2
			if i == j {
				continue
			}
			for k = 1; k <= 3; k++ { // k = column number to use for row #3
				if i == k || j == k {
					continue
				}
				P.M[0][0] = 0.0
				P.M[0][1] = 0.0
				P.M[0][2] = 0.0
				P.M[1][0] = 0.0
				P.M[1][1] = 0.0
				P.M[1][2] = 0.0
				P.M[2][0] = 0.0
				P.M[2][1] = 0.0
				P.M[2][2] = 0.0
				for p = -1; p <= 1; p += 2 { // p,q,r are -1 or +1
					for q = -1; q <= 1; q += 2 { // and go into rows #1,2,3
						for r = -1; r <= 1; r += 2 {
							P.M[0][i-1] = float64(p)
							P.M[1][j-1] = float64(q)
							P.M[2][k-1] = float64(r)
							detP = matrix.Mat33Determinant(P) // sign of permutation
							if detP*detQ <= 0.0 {
								continue
							} // doesn't match sign of Q
							M = matrix.MatMultiply(P, Q)

							// angle of M rotation = 2.0*acos(0.5*sqrt(1.0+trace(M)))
							// we want largest trace(M) == smallest angle == M nearest to I

							val = M.M[0][0] + M.M[1][1] + M.M[2][2] /* trace */
							if val > vbest {
								vbest = val
								ibest = i
								jbest = j
								kbest = k
								pbest = p
								qbest = q
								rbest = r
							}
						}
					}
				}
			}
		}
	}

	switch ibest * pbest {
	case 1:
		i = constant.NIFTI_L2R
	case -1:
		i = constant.NIFTI_R2L
	case 2:
		i = constant.NIFTI_P2A
	case -2:
		i = constant.NIFTI_A2P
	case 3:
		i = constant.NIFTI_I2S
	case -3:
		i = constant.NIFTI_S2I
	default:
		i = 0
	}

	switch jbest * qbest {
	case 1:
		j = constant.NIFTI_L2R
	case -1:
		j = constant.NIFTI_R2L
	case 2:
		j = constant.NIFTI_P2A
	case -2:
		j = constant.NIFTI_A2P
	case 3:
		j = constant.NIFTI_I2S
	case -3:
		j = constant.NIFTI_S2I
	default:
		j = 0
	}

	switch kbest * rbest {
	case 1:
		k = constant.NIFTI_L2R
	case -1:
		k = constant.NIFTI_R2L
	case 2:
		k = constant.NIFTI_P2A
	case -2:
		k = constant.NIFTI_A2P
	case 3:
		k = constant.NIFTI_I2S
	case -3:
		k = constant.NIFTI_S2I
	default:
		k = 0
	}

	res := [3]int32{i, j, k}
	n.Data.IJKOrtient = res

}
