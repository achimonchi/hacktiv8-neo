package pkg

/*
	Token :
		- userId
		- expired
	1. GenereateToken(payload *Token)(string, error)
	2. VerifyToken(token string)(*Token, error)

	Unit Test
	1. TestGenerateToken_Success
	2. TestGenerateToken_PayloadNull_Fail
	3. VerifyToken_Success
	4. VerifyToken_InvalidToken
	5. VerifyToken_Expired
*/
