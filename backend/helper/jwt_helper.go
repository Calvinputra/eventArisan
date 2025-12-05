package helper

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt"
)

func GenerateRandomKey(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func GetRecidFromJWTC(ctx *gin.Context) (string, error) {
	tokenRequest := ctx.Request.Header.Get("Authorization")

	re := regexp.MustCompile(`^Bearer\s.+$`)
	if tokenRequest == "" || !re.MatchString(tokenRequest) {
		return "", fmt.Errorf("JWT is not in the right format: %s", tokenRequest)
	}

	auth := strings.Split(tokenRequest, ` `)
	token, _, err := new(jwt.Parser).ParseUnverified(auth[1], jwt.MapClaims{})
	if err != nil {
		return "", fmt.Errorf("failed to parse the token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println("Claim", claims)
		inputter := fmt.Sprint(claims["recid"])
		if inputter == "" || inputter == "<nil>" {
			return "", fmt.Errorf("invalid inputter value")
		}
		return inputter, nil
	}

	return "", fmt.Errorf("failed to map JWT claim")
}

func GetEmail(ctx *gin.Context) (string, error) {
	tokenString := ctx.Request.Header.Get("Authorization")
	re := regexp.MustCompile(`^Bearer\s.+$`)
	if tokenString == "" || !re.MatchString(tokenString) {
		fmt.Println("Empty authorization format is invalid")
		return "", errors.New("invalid authorization format")
	}
	auth := strings.Split(tokenString, " ")
	token, _, err := new(jwt.Parser).ParseUnverified(auth[1], jwt.MapClaims{})
	if err != nil {
		fmt.Println("Empty %s", err)
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Empty: Couldn't extract claims")
		return "", err
	}
	emailInterface, ok := claims["email"]
	if !ok {
		fmt.Println("Empty: Couldn't extract email")
		return "", err
	}
	return emailInterface.(string), nil
}

func GetBranchCodesFromJWT(ctx *gin.Context) []string {
	tokenString := ctx.Request.Header.Get("Authorization")

	re := regexp.MustCompile(`^Bearer\s.+$`)
	if tokenString == "" || !re.MatchString(tokenString) {
		fmt.Println("Empty authorization format is invalid")
		return nil
	}
	auth := strings.Split(tokenString, " ")
	// log.Printf("token string %s", auth[1])
	token, _, err := new(jwt.Parser).ParseUnverified(auth[1], jwt.MapClaims{})
	if err != nil {
		fmt.Println("Empty %s", err)
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Empty: Couldn't extract claims")
		return nil
	}

	// TypeRecid assertion to get the value as a slice of interfaces
	branchCodeInterface, ok := claims["branchCodes"].([]interface{})
	if !ok {
		fmt.Println("Empty: Couldn't extract branchCodes as []interface{}")
		return nil
	}

	// Convert the slice of interfaces to a slice of strings
	var branchCode []string
	for _, recid := range branchCodeInterface {
		if recidStr, ok := recid.(string); ok {
			branchCode = append(branchCode, recidStr)
		} else {
			fmt.Println("Empty: Element in branchCode is not a string")
			return nil
		}
	}

	return branchCode
}
