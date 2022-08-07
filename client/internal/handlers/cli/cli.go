package cli

import (
	"bufio"
	"client/internal/application"
	"client/internal/config"
	"client/internal/domain"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
)

type ConsoleCommandHandler struct {
	config          config.Config
	userService     application.UserService
	userDataService application.UserDataService
}

func GetConsoleCommandHandler(
	config config.Config,
	userService application.UserService,
	userDataService application.UserDataService,
) ConsoleCommandHandler {
	return ConsoleCommandHandler{
		config:          config,
		userService:     userService,
		userDataService: userDataService,
	}
}

func (h ConsoleCommandHandler) RunDialog() {
	if h.config.GetToken() == "" {
		h.auth()
	} else {
		h.enterMainCommand()
	}
}

func (h ConsoleCommandHandler) auth() {
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{"register", "login", "quit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "register":
		h.register()
	case "login":
		h.login()
	case "quit":
		h.quit()
	default:
		h.auth()
	}

	fmt.Printf("You choose %q\n", result)
}

func (h ConsoleCommandHandler) register() {
	fmt.Printf("Enter your email:")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	fmt.Printf("Enter your password:")
	reader = bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	fmt.Printf("Repeat your password:")
	reader = bufio.NewReader(os.Stdin)
	repeatPassword, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	token, err := h.userService.Register(
		domain.User{
			Email:          email,
			Password:       password,
			RepeatPassword: repeatPassword,
		},
	)

	//////////////////
	log.Println(token) //ToDo: delete it
	log.Println(err)   //ToDo: delete it

	if err != nil || token == "" {
		fmt.Print("registration error. ")

		if err != nil {
			fmt.Println(err)
		}

		h.RunDialog()
	}

	h.config.SetToken(token)
	h.RunDialog()
}

func (h ConsoleCommandHandler) login() {
	fmt.Printf("Enter your email:")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	fmt.Printf("Enter your password:")
	reader = bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	token, err := h.userService.Auth(
		domain.User{Email: email, Password: password},
	)

	if err != nil || token == "" {
		fmt.Print("login error. ")

		if err != nil {
			fmt.Println(err)
		}

		h.RunDialog()
	}

	h.config.SetToken(token)
	h.RunDialog()
}

func (h ConsoleCommandHandler) quit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func (h ConsoleCommandHandler) enterMainCommand() {
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{
			"get-all-data",
			"get-login-password",
			"get-credit-card",
			"get-text-record",
			"get-binary-record",
			"quit",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "get-all-data":
		userData, err := h.userDataService.GetAllData()
		fmt.Println(userData)
		fmt.Println(err)
	case "get-login-password":
		loginPasswordCollection, err := h.userDataService.GetLoginPasswordCollection()
		fmt.Println(loginPasswordCollection)
		fmt.Println(err)
	case "get-credit-card":
		creditCardCollection, err := h.userDataService.GetCreditCardCollection()
		fmt.Println(creditCardCollection)
		fmt.Println(err)
	case "get-text-record":
		textRecordCollection, err := h.userDataService.GetTextRecordCollection()
		fmt.Println(textRecordCollection)
		fmt.Println(err)
	case "get-binary-record":
		binaryRecordCollection, err := h.userDataService.GetBinaryRecordCollection()
		fmt.Println(binaryRecordCollection)
		fmt.Println(err)
	case "quit":
		h.quit()
	default:
		h.enterMainCommand()
	}
}
