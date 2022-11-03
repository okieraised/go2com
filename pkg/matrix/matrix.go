package matrix

import "C"
import "math"

type FMat44 struct {
	M [4][4]float32
}

type FMat33 struct {
	M [3][3]float32
}

type DMat44 struct {
	M [4][4]float64
}

type DMat33 struct {
	M [3][3]float64
}

// Mat33Polar finds the closest orthogonal matrix to input A
// (in both the Frobenius and L2 norms).
// Algorithm is that from NJ Higham, SIAM J Sci Stat Comput, 7:1160-1174.
func Mat33Polar(A DMat33) DMat33 {
	var X, Y, Z DMat33
	var alp, bet, gam, gmi float64
	var dif float64 = 1.0
	var k int64 = 0

	X = A

	gam = Mat33Determinant(X)
	for {
		if gam > 0.0 {
			break
		}
		gam = 0.00001 * (0.001 + Mat33RowNorm(X))
		X.M[0][0] += gam
		X.M[1][1] += gam
		X.M[2][2] += gam
		gam = Mat33Determinant(X)
	}

	for {
		Y = Mat33Inverse(X)
		if dif > 0.3 { /* far from convergence */
			alp = math.Sqrt(Mat33RowNorm(X) * Mat33ColNorm(X))
			bet = math.Sqrt(Mat33RowNorm(Y) * Mat33ColNorm(Y))
			gam = math.Sqrt(bet / alp)
			gmi = 1.0 / gam
		} else {
			gam = 1.0
			gmi = 1.0
		}
		Z.M[0][0] = 0.5 * (gam*X.M[0][0] + gmi*Y.M[0][0])
		Z.M[0][1] = 0.5 * (gam*X.M[0][1] + gmi*Y.M[1][0])
		Z.M[0][2] = 0.5 * (gam*X.M[0][2] + gmi*Y.M[2][0])
		Z.M[1][0] = 0.5 * (gam*X.M[1][0] + gmi*Y.M[0][1])
		Z.M[1][1] = 0.5 * (gam*X.M[1][1] + gmi*Y.M[1][1])
		Z.M[1][2] = 0.5 * (gam*X.M[1][2] + gmi*Y.M[2][1])
		Z.M[2][0] = 0.5 * (gam*X.M[2][0] + gmi*Y.M[0][2])
		Z.M[2][1] = 0.5 * (gam*X.M[2][1] + gmi*Y.M[1][2])
		Z.M[2][2] = 0.5 * (gam*X.M[2][2] + gmi*Y.M[2][2])

		dif = math.Abs(Z.M[0][0]-X.M[0][0]) + math.Abs(Z.M[0][1]-X.M[0][1]) + math.Abs(Z.M[0][2]-X.M[0][2]) +
			math.Abs(Z.M[1][0]-X.M[1][0]) + math.Abs(Z.M[1][1]-X.M[1][1]) + math.Abs(Z.M[1][2]-X.M[1][2]) +
			math.Abs(Z.M[2][0]-X.M[2][0]) + math.Abs(Z.M[2][1]-X.M[2][1]) + math.Abs(Z.M[2][2]-X.M[2][2])

		k = k + 1
		if k > 100 || dif < 3.e-6 { // convergence or exhaustion
			break
		}
		X = Z
	}
	return Z
}

// Mat33Determinant computes the determinant of a 3x3 matrix
func Mat33Determinant(R DMat33) float64 {
	var r11, r12, r13, r21, r22, r23, r31, r32, r33 float64

	// [ r11 r12 r13 ]
	// [ r21 r22 r23 ]
	// [ r31 r32 r33 ]

	r11 = R.M[0][0]
	r12 = R.M[0][1]
	r13 = R.M[0][2]
	r21 = R.M[1][0]
	r22 = R.M[1][1]
	r23 = R.M[1][2]
	r31 = R.M[2][0]
	r32 = R.M[2][1]
	r33 = R.M[2][2]

	return r11*r22*r33 - r11*r32*r23 - r21*r12*r33 + r21*r32*r13 + r31*r12*r23 - r31*r22*r13
}

// Mat33RowNorm computes the max row norm of a 3x3 matrix
func Mat33RowNorm(A DMat33) float64 {
	var r1, r2, r3 float64

	r1 = math.Abs(A.M[0][0]) + math.Abs(A.M[0][1]) + math.Abs(A.M[0][2])
	r2 = math.Abs(A.M[1][0]) + math.Abs(A.M[1][1]) + math.Abs(A.M[1][2])
	r3 = math.Abs(A.M[2][0]) + math.Abs(A.M[2][1]) + math.Abs(A.M[2][2])
	if r1 < r2 {
		r1 = r2
	}
	if r1 < r3 {
		r1 = r3
	}
	return r1
}

// Mat33ColNorm computes the max column norm of a 3x3 matrix
func Mat33ColNorm(A DMat33) float64 {
	var r1, r2, r3 float64

	r1 = math.Abs(A.M[0][0]) + math.Abs(A.M[1][0]) + math.Abs(A.M[2][0])
	r2 = math.Abs(A.M[0][1]) + math.Abs(A.M[1][1]) + math.Abs(A.M[2][1])
	r3 = math.Abs(A.M[0][2]) + math.Abs(A.M[1][2]) + math.Abs(A.M[2][2])

	if r1 < r2 {
		r1 = r2
	}
	if r1 < r3 {
		r1 = r3
	}
	return r1
}

// MatMultiply multiples 2 3x3 matrices
func MatMultiply(A, B DMat33) DMat33 {
	var C DMat33
	var i, j int64

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			C.M[i][j] = A.M[i][0]*B.M[0][j] + A.M[i][1]*B.M[1][j] + A.M[i][2]*B.M[2][j]
		}
	}
	return C
}

// Mat33Inverse computes the inverse of a bordered 3x3 matrix
func Mat33Inverse(R DMat33) DMat33 {
	var r11, r12, r13, r21, r22, r23, r31, r32, r33, deti float64
	var Q DMat33

	//  INPUT MATRIX:
	// [ r11 r12 r13 ]
	// [ r21 r22 r23 ]
	// [ r31 r32 r33 ]

	r11 = R.M[0][0]
	r12 = R.M[0][1]
	r13 = R.M[0][2]
	r21 = R.M[1][0]
	r22 = R.M[1][1]
	r23 = R.M[1][2]
	r31 = R.M[2][0]
	r32 = R.M[2][1]
	r33 = R.M[2][2]

	deti = r11*r22*r33 - r11*r32*r23 - r21*r12*r33 + r21*r32*r13 + r31*r12*r23 - r31*r22*r13

	if deti != 0.0 {
		deti = 1.0 / deti
	}

	Q.M[0][0] = deti * (r22*r33 - r32*r23)
	Q.M[0][1] = deti * (-r12*r33 + r32*r13)
	Q.M[0][2] = deti * (r12*r23 - r22*r13)

	Q.M[1][0] = deti * (-r21*r33 + r31*r23)
	Q.M[1][1] = deti * (r11*r33 - r31*r13)
	Q.M[1][2] = deti * (-r11*r23 + r21*r13)

	Q.M[2][0] = deti * (r21*r32 - r31*r22)
	Q.M[2][1] = deti * (-r11*r32 + r31*r12)
	Q.M[2][2] = deti * (r11*r22 - r21*r12)

	return Q

}

// Mat44Inverse computes the inverse of a bordered 4x4 matrix
func Mat44Inverse(R DMat44) DMat44 {
	var r11, r12, r13, r21, r22, r23, r31, r32, r33, v1, v2, v3, deti float64
	var Q DMat44

	//  INPUT MATRIX IS
	// [ r11 r12 r13 v1 ]
	// [ r21 r22 r23 v2 ]
	// [ r31 r32 r33 v3 ]
	// [  0   0   0   1 ]
	r11 = R.M[0][0]
	r12 = R.M[0][1]
	r13 = R.M[0][2]
	r21 = R.M[1][0]
	r22 = R.M[1][1]
	r23 = R.M[1][2]
	r31 = R.M[2][0]
	r32 = R.M[2][1]
	r33 = R.M[2][2]
	v1 = R.M[0][3]
	v2 = R.M[1][3]
	v3 = R.M[2][3]

	deti = r11*r22*r33 - r11*r32*r23 - r21*r12*r33 + r21*r32*r13 + r31*r12*r23 - r31*r22*r13 // determinant

	if deti != 0.0 {
		deti = 1.0 / deti
	}

	Q.M[0][0] = deti * (r22*r33 - r32*r23)
	Q.M[0][1] = deti * (-r12*r33 + r32*r13)
	Q.M[0][2] = deti * (r12*r23 - r22*r13)
	Q.M[0][3] = deti * (-r12*r23*v3 + r12*v2*r33 + r22*r13*v3 - r22*v1*r33 - r32*r13*v2 + r32*v1*r23)

	Q.M[1][0] = deti * (-r21*r33 + r31*r23)
	Q.M[1][1] = deti * (r11*r33 - r31*r13)
	Q.M[1][2] = deti * (-r11*r23 + r21*r13)
	Q.M[1][3] = deti * (r11*r23*v3 - r11*v2*r33 - r21*r13*v3 + r21*v1*r33 + r31*r13*v2 - r31*v1*r23)

	Q.M[2][0] = deti * (r21*r32 - r31*r22)
	Q.M[2][1] = deti * (-r11*r32 + r31*r12)
	Q.M[2][2] = deti * (r11*r22 - r21*r12)
	Q.M[2][3] = deti * (-r11*r22*v3 + r11*r32*v2 + r21*r12*v3 - r21*r32*v1 - r31*r12*v2 + r31*r22*v1)

	Q.M[3][0] = 0.0
	Q.M[3][1] = 0.0
	Q.M[3][2] = 0.0
	if deti == 0.0 {
		Q.M[3][3] = 0.0
	} else {
		Q.M[3][3] = 1.0
	}

	return Q

}

// MakeOrthoMat44 uses 9 input float64 and make an orthogonal mat44 out of them
func MakeOrthoMat44(r11, r12, r13, r21, r22, r23, r31, r32, r33 float64) DMat44 {
	var R DMat44
	var Q, P DMat33
	var val float64

	R.M[3][0] = 0.0
	R.M[3][1] = 0.0
	R.M[3][2] = 0.0
	R.M[3][3] = 1.0

	Q.M[0][0] = r11
	Q.M[0][1] = r12
	Q.M[0][2] = r13
	Q.M[1][0] = r21
	Q.M[1][1] = r22
	Q.M[1][2] = r23
	Q.M[2][0] = r31
	Q.M[2][1] = r32
	Q.M[2][2] = r33

	// normalize row 1
	val = Q.M[0][0]*Q.M[0][0] + Q.M[0][1]*Q.M[0][1] + Q.M[0][2]*Q.M[0][2]
	if val > 0.0 {
		val = 1.0 / math.Sqrt(val)
		Q.M[0][0] *= val
		Q.M[0][1] *= val
		Q.M[0][2] *= val
	} else {
		Q.M[0][0] = 1.0
		Q.M[0][1] = 0.0
		Q.M[0][2] = 0.0
	}

	// normalize row 2
	val = Q.M[1][0]*Q.M[1][0] + Q.M[1][1]*Q.M[1][1] + Q.M[1][2]*Q.M[1][2]
	if val > 0.0 {
		val = 1.0 / math.Sqrt(val)
		Q.M[1][0] *= val
		Q.M[1][1] *= val
		Q.M[1][2] *= val
	} else {
		Q.M[1][0] = 0.0
		Q.M[1][1] = 1.0
		Q.M[1][2] = 0.0
	}

	// normalize row 3
	val = Q.M[2][0]*Q.M[2][0] + Q.M[2][1]*Q.M[2][1] + Q.M[2][2]*Q.M[2][2]
	if val > 0.0 {
		val = 1.0 / math.Sqrt(val)
		Q.M[2][0] *= val
		Q.M[2][1] *= val
		Q.M[2][2] *= val
	} else {
		Q.M[2][0] = Q.M[0][1]*Q.M[1][2] - Q.M[0][2]*Q.M[1][1]
		Q.M[2][1] = Q.M[0][2]*Q.M[1][0] - Q.M[0][0]*Q.M[1][2]
		Q.M[2][2] = Q.M[0][0]*Q.M[1][1] - Q.M[0][1]*Q.M[1][0]
	}

	// P is orthogonal matrix closest to Q
	P = Mat33Polar(Q)

	R.M[0][0] = P.M[0][0]
	R.M[0][1] = P.M[0][1]
	R.M[0][2] = P.M[0][2]
	R.M[1][0] = P.M[1][0]
	R.M[1][1] = P.M[1][1]
	R.M[1][2] = P.M[1][2]
	R.M[2][0] = P.M[2][0]
	R.M[2][1] = P.M[2][1]
	R.M[2][2] = P.M[2][2]

	R.M[0][3] = 0.0
	R.M[1][3] = 0.0
	R.M[2][3] = 0.0

	return R
}
