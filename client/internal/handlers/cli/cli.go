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
	"strings"
	"text/tabwriter"
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
}

func (h ConsoleCommandHandler) getItemType() string {
	prompt := promptui.Select{
		Label: "Choose type for getting detail info",
		Items: []string{"login password", "credit card", "text record", "binary record", "back"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func (h ConsoleCommandHandler) register() {
	fmt.Printf("Enter your email:")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	email = strings.TrimSuffix(email, "\n")

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	fmt.Printf("Enter your password:")
	reader = bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	fmt.Printf("Repeat your password:")
	reader = bufio.NewReader(os.Stdin)
	repeatPassword, err := reader.ReadString('\n')
	repeatPassword = strings.TrimSuffix(repeatPassword, "\n")

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
	email = strings.TrimSuffix(email, "\n")

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	fmt.Printf("Enter your password:")
	reader = bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

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
			"create-new-item",
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
		h.getAllData()
	case "get-login-password":
		loginPasswordCollection, err := h.userDataService.GetLoginPasswordCollection(h.config.GetToken())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		h.showList("Login passwords", loginPasswordCollection)
		h.getLoginPassword(loginPasswordCollection)
	case "get-credit-card":
		creditCardCollection, err := h.userDataService.GetCreditCardCollection(h.config.GetToken())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		h.showList("Credit cards", creditCardCollection)
		h.getCreditCard(creditCardCollection)
	case "get-text-record":
		textRecordCollection, err := h.userDataService.GetTextRecordCollection(h.config.GetToken())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		h.showList("Text records", textRecordCollection)
		h.getTextRecord(textRecordCollection)
	case "get-binary-record":
		binaryRecordCollection, err := h.userDataService.GetBinaryRecordCollection(h.config.GetToken())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		h.showList("Binary records", binaryRecordCollection)
		h.getBinaryRecord(binaryRecordCollection)
	case "create-new-item":
		h.createNewItem()
	case "quit":
		h.quit()
	default:
		h.enterMainCommand()
	}

	h.RunDialog()
}

func (h ConsoleCommandHandler) getAllData() {
	userData, err := h.userDataService.GetAllData(h.config.GetToken())

	if err != nil {
		fmt.Println(err)
	}

	h.showList("Login passwords", userData.LoginPasswordCollection)
	h.showList("Credit cards", userData.CreditCardCollection)
	h.showList("Binary records", userData.BinaryRecordCollection)
	h.showList("Text records", userData.TextRecordCollection)

	itemType := h.getItemType()

	switch itemType {
	case "login password":
		h.getLoginPassword(userData.LoginPasswordCollection)
	case "credit card":
		h.getCreditCard(userData.CreditCardCollection)
	case "binary record":
		h.getBinaryRecord(userData.BinaryRecordCollection)
	case "text record":
		h.getBinaryRecord(userData.TextRecordCollection)
	case "create-new-item":
		h.getBinaryRecord(userData.TextRecordCollection)
	case "back":
		h.RunDialog()
	default:
		h.RunDialog()
	}

	h.RunDialog()
}

func (h ConsoleCommandHandler) createNewItem() {
	itemType := h.getItemType()

	switch itemType {
	case "login password":
		item, err := h.createLoginPassword()

		if err != nil {
			log.Println(err)
		}

		h.showLoginPassword(*item)
	case "credit card":
		item, err := h.createCreditCard()

		if err != nil {
			log.Println(err)
		}

		h.showCreditCard(*item)
	case "binary record":
		item, err := h.createBinaryRecord()

		if err != nil {
			log.Println(err)
		}

		h.showBinaryRecord(*item)
	case "text record":
		item, err := h.createTextRecord()

		if err != nil {
			log.Println(err)
		}

		h.showTextRecord(*item)
	case "back":
		h.RunDialog()
	default:
		h.RunDialog()
	}

	h.RunDialog()
}

func (h ConsoleCommandHandler) createLoginPassword() (*domain.LoginPassword, error) {
	loginPassword := h.getLoginPasswordFields()

	item, err := h.userDataService.CreateLoginPassword(h.config.GetToken(), loginPassword)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getLoginPasswordFields() domain.LoginPassword {
	loginPassword := domain.LoginPassword{}

	fmt.Printf("Enter name:")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	loginPassword.Name = strings.TrimSuffix(name, "\n")

	fmt.Printf("Enter login:")
	reader = bufio.NewReader(os.Stdin)
	login, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	loginPassword.Login = strings.TrimSuffix(login, "\n")

	fmt.Printf("Enter password:")
	reader = bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	loginPassword.Password = strings.TrimSuffix(password, "\n")

	fmt.Printf("Enter note:")
	reader = bufio.NewReader(os.Stdin)
	note, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	loginPassword.Note = strings.TrimSuffix(note, "\n")

	return loginPassword
}

func (h ConsoleCommandHandler) createCreditCard() (*domain.CreditCard, error) {
	creditCard := h.getCreditCardFields()

	item, err := h.userDataService.CreateCreditCard(h.config.GetToken(), creditCard)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getCreditCardFields() domain.CreditCard {
	creditCard := domain.CreditCard{}

	fmt.Printf("Enter name:")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	creditCard.Name = strings.TrimSuffix(name, "\n")

	fmt.Printf("Enter number:")
	reader = bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	creditCard.Number = strings.TrimSuffix(number, "\n")

	fmt.Printf("Enter expiration:")
	reader = bufio.NewReader(os.Stdin)
	expiration, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	creditCard.Expiration = strings.TrimSuffix(expiration, "\n")

	fmt.Printf("Enter cvv:")
	reader = bufio.NewReader(os.Stdin)
	cvv, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	creditCard.Cvv = strings.TrimSuffix(cvv, "\n")

	fmt.Printf("Enter owner:")
	reader = bufio.NewReader(os.Stdin)
	owner, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	creditCard.Cvv = strings.TrimSuffix(owner, "\n")

	fmt.Printf("Enter note:")
	reader = bufio.NewReader(os.Stdin)
	note, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	creditCard.Note = strings.TrimSuffix(note, "\n")

	return creditCard
}

func (h ConsoleCommandHandler) createBinaryRecord() (*domain.BinaryRecord, error) {
	binaryRecord := h.getBinaryRecordFields()

	item, err := h.userDataService.CreateBinaryRecord(h.config.GetToken(), binaryRecord)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getBinaryRecordFields() domain.BinaryRecord {
	binaryRecord := domain.BinaryRecord{}

	fmt.Printf("Enter name:")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	binaryRecord.Name = strings.TrimSuffix(name, "\n")

	fmt.Printf("Enter binary data:")
	reader = bufio.NewReader(os.Stdin)
	binary, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	binaryRecord.Binary = strings.TrimSuffix(binary, "\n")

	fmt.Printf("Enter note:")
	reader = bufio.NewReader(os.Stdin)
	note, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	binaryRecord.Note = strings.TrimSuffix(note, "\n")

	return binaryRecord
}

func (h ConsoleCommandHandler) createTextRecord() (*domain.TextRecord, error) {
	textRecord := h.getTextRecordFields()

	item, err := h.userDataService.CreateTextRecord(h.config.GetToken(), textRecord)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getTextRecordFields() domain.TextRecord {
	textRecord := domain.TextRecord{}

	fmt.Printf("Enter name:")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	textRecord.Name = strings.TrimSuffix(name, "\n")

	fmt.Printf("Enter text data:")
	reader = bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	textRecord.Text = strings.TrimSuffix(text, "\n")

	fmt.Printf("Enter note:")
	reader = bufio.NewReader(os.Stdin)
	note, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.RunDialog()
	}

	textRecord.Note = strings.TrimSuffix(note, "\n")

	return textRecord
}

func (h ConsoleCommandHandler) getLoginPassword(collection []domain.ProtectedItem) {
	if len(collection) == 0 {
		h.RunDialog()
	}

	itemId := h.getItemIdByCollection(collection)

	if itemId != "" {
		loginPassword, err := h.getLoginPasswordById(itemId)

		if err != nil {
			log.Println(err)
			return
		}

		h.showLoginPassword(*loginPassword)
	}
}

func (h ConsoleCommandHandler) getCreditCard(collection []domain.ProtectedItem) {
	if len(collection) == 0 {
		h.RunDialog()
	}

	itemId := h.getItemIdByCollection(collection)

	if itemId != "" {
		creditCard, err := h.getCreditCardById(itemId)

		if err != nil {
			log.Println(err)
			return
		}

		h.showCreditCard(*creditCard)
	}
}

func (h ConsoleCommandHandler) getTextRecord(collection []domain.ProtectedItem) {
	if len(collection) == 0 {
		h.RunDialog()
	}

	itemId := h.getItemIdByCollection(collection)

	if itemId != "" {
		text, err := h.getTextById(itemId)

		if err != nil {
			log.Println(err)
			return
		}

		h.showTextRecord(*text)
	}
}

func (h ConsoleCommandHandler) getBinaryRecord(collection []domain.ProtectedItem) {
	if len(collection) == 0 {
		h.RunDialog()
	}

	itemId := h.getItemIdByCollection(collection)

	if itemId != "" {
		binary, err := h.getBinaryById(itemId)

		if err != nil {
			log.Println(err)
			return
		}

		h.showBinaryRecord(*binary)
	}
}

func (h ConsoleCommandHandler) getLoginPasswordById(id string) (*domain.LoginPassword, error) {
	item, err := h.userDataService.GetLoginPasswordById(h.config.GetToken(), id)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getCreditCardById(id string) (*domain.CreditCard, error) {
	item, err := h.userDataService.GetCreditCardById(h.config.GetToken(), id)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getBinaryById(id string) (*domain.BinaryRecord, error) {
	item, err := h.userDataService.GetBinaryRecordById(h.config.GetToken(), id)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getTextById(id string) (*domain.TextRecord, error) {
	item, err := h.userDataService.GetTextRecordById(h.config.GetToken(), id)

	if err != nil {
		return nil, err
	}

	return item, err
}

func (h ConsoleCommandHandler) getItemIdByCollection(collections []domain.ProtectedItem) string {
	var idCollection []string

	for _, item := range collections {
		idCollection = append(idCollection, item.Id)
	}

	prompt := promptui.Select{
		Label: "Choose item id for getting detail info",
		Items: idCollection,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func (h ConsoleCommandHandler) showList(name string, list []domain.ProtectedItem) error {
	fmt.Println("####################################")
	fmt.Println(name)
	fmt.Println("####################################")
	w := tabwriter.NewWriter(os.Stdout, 5, 1, 1, ' ', 0)
	fmt.Fprintln(w, "id\tname\tcreated_at\tlast_access_at\t")

	if len(list) == 0 {
		fmt.Fprintln(w, "no data...\t")
	}

	for _, item := range list {
		fmt.Fprintln(w, item.Id+"\t"+item.Name+"\t"+item.CreatedAt+"\t"+item.LastAccessAt)
	}

	err := w.Flush()

	return err
}

func (h ConsoleCommandHandler) showLoginPassword(item domain.LoginPassword) error {
	fmt.Println("####################################")
	fmt.Println("Detail info about login password with id=" + item.Id)
	fmt.Println("####################################")
	w := tabwriter.NewWriter(os.Stdout, 5, 1, 1, ' ', 0)
	fmt.Fprintln(w, "id\tname\tlogin\tpassword\tnote\tcreated_at\tlast_access_at\t")
	fmt.Fprintln(
		w,
		item.Id+"\t"+item.Name+"\t"+item.Login+"\t"+item.Password+
			"\t"+item.Note+"\t"+item.CreatedAt+"\t"+item.LastAccessAt,
	)

	err := w.Flush()

	return err
}

func (h ConsoleCommandHandler) showCreditCard(item domain.CreditCard) error {
	fmt.Println("####################################")
	fmt.Println("Detail info about credit card with id=" + item.Id)
	fmt.Println("####################################")
	w := tabwriter.NewWriter(os.Stdout, 5, 1, 1, ' ', 0)
	fmt.Fprintln(w, "id\tname\tnumber\texp\tcvv\towner\tnote\tcreated_at\tlast_access_at\t")
	fmt.Fprintln(
		w,
		item.Id+"\t"+
			item.Name+"\t"+
			item.Number+"\t"+
			item.Expiration+"\t"+
			item.Cvv+"\t"+
			item.Owner+"\t"+
			item.Note+"\t"+
			item.CreatedAt+"\t"+
			item.LastAccessAt,
	)

	err := w.Flush()

	return err
}

func (h ConsoleCommandHandler) showTextRecord(item domain.TextRecord) error {
	fmt.Println("####################################")
	fmt.Println("Detail info about text record with id=" + item.Id)
	fmt.Println("####################################")

	w := tabwriter.NewWriter(os.Stdout, 5, 1, 1, ' ', 0)
	fmt.Fprintln(w, "id\tname\ttext\tnote\tcreated_at\tlast_access_at\t")
	fmt.Fprintln(
		w,
		item.Id+"\t"+
			item.Name+"\t"+
			item.Text+"\t"+
			item.Note+"\t"+
			item.CreatedAt+"\t"+
			item.LastAccessAt,
	)

	err := w.Flush()

	return err
}

func (h ConsoleCommandHandler) showBinaryRecord(item domain.BinaryRecord) error {
	fmt.Println("####################################")
	fmt.Println("Detail info about binary record with id=" + item.Id)
	fmt.Println("####################################")

	w := tabwriter.NewWriter(os.Stdout, 5, 1, 1, ' ', 0)
	fmt.Fprintln(w, "id\tname\tbinary\tnote\tcreated_at\tlast_access_at\t")
	fmt.Fprintln(
		w,
		item.Id+"\t"+
			item.Name+"\t"+
			item.Binary+"\t"+
			item.Note+"\t"+
			item.CreatedAt+"\t"+
			item.LastAccessAt,
	)

	err := w.Flush()

	return err
}
