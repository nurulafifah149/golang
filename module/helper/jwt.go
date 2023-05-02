package helper

import (
	"time"

	"github.com/kataras/jwt"
	"github.com/nurulafifah149/golang/module/model/token"
	"github.com/nurulafifah149/golang/module/model/user"
	MyLog "github.com/nurulafifah149/golang/pkg/logger"
)

var SharedKey = []byte("sercrethatmaycontainch@r$32charS")

type Claims struct {
	AccessClaims  token.AccessClaim
	DefaultClaims token.DefaultClaim
}

func GenerateDefaultClaims(username string) token.DefaultClaim {
	timenow := time.Now()
	return token.DefaultClaim{
		Expired:   int(timenow.Add(24 * time.Hour).UnixMilli()),
		NotBefore: int(timenow.UnixMilli()),
		IssuedAt:  int(timenow.UnixMilli()),
		Issuer:    "myGram",
		Audience:  "golang",
		JTI:       username,
		Typ:       "",
	}
}

func GenerateToken(userInfo user.User) (tokenOut token.Tokens, err error) {
	MyLog.LogMyApp("i", "Generate JWT Invoked", "Helper - GenerateToken", nil)
	MyLog.LogMyApp("i", "Starting Generate Default Claim", "Helper - GenerateToken", nil)
	defaultClaim := GenerateDefaultClaims(userInfo.Username)
	defaultClaim.Typ = "id_token"

	MyLog.LogMyApp("i", "Starting Generate ID Claim", "Helper - GenerateToken", nil)
	AccessClaim := token.AccessClaim{
		UserId:   userInfo.Id,
		Username: userInfo.Username,
		Role:     userInfo.Role,
	}

	userClaims := Claims{
		AccessClaims:  AccessClaim,
		DefaultClaims: defaultClaim,
	}

	//generate JWT
	MyLog.LogMyApp("i", "Starting Generate Id Token", "Helper - GenerateToken", nil)
	Idtoken, err := jwt.Sign(jwt.HS256, SharedKey, userClaims)
	tokenOut.AccessToken = string(Idtoken)
	if err != nil {
		MyLog.LogMyApp("e", "Eror When Sign a Token", "Helper - GenerateToken", nil)
	}

	return
}

func VerifyAndParse(token string, Claims any) (err error) {
	MyLog.LogMyApp("i", "Starting Verify and Parse Claim", "Helper - GenerateToken", nil)
	verifiedToken, err := jwt.Verify(jwt.HS256, SharedKey, []byte(token))
	if err != nil {
		MyLog.LogMyApp("e", "Eror When Verify Token", "Helper - GenerateToken", nil)
		return
	}

	err = verifiedToken.Claims(&Claims)
	if err != nil {
		MyLog.LogMyApp("e", "Eror When Parse Token", "Helper - GenerateToken", nil)
	}
	return
}
