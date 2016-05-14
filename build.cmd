set GOPATH=f:\work\saog\authserver
set PATH=%PATH%;f:\work\saog\authserver\bin

go get github.com/tools/godep

godep get github.com/BurntSushi/toml github.com/gorilla/mux github.com/SermoDigital/jose github.com/tamnd/gauth github.com/tamnd/httpclient golang.org/x/oauth2 github.com/gokrokvertskhov/gauth github.com/dgrijalva/jwt-go github.com/gokrokvertskhov/serverlib github.com/abbot/go-http-auth

godep save github.com/BurntSushi/toml github.com/gorilla/mux github.com/SermoDigital/jose github.com/tamnd/gauth github.com/tamnd/httpclient golang.org/x/oauth2 github.com/gokrokvertskhov/gauth github.com/dgrijalva/jwt-go github.com/gokrokvertskhov/serverlib github.com/abbot/go-http-auth
