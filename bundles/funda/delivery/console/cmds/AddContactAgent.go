package cmds

import (
	"fmt"
	"github.com/fspace/ecm/bundles/funda/gateways/memory"
	"github.com/fspace/ecm/bundles/funda/usecases"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Context for "ls" command
type AddContactAgent struct {
	//All bool
	// Data string
	usecases.ContactAgentRequestMessage
}

func (cmd *AddContactAgent) Run(c *kingpin.ParseContext) error {
	//fmt.Printf("all=%v\n", l.All)
	fmt.Println("AddContactAgent -->")
	fmt.Println("AddContactAgent value :", cmd.ContactAgentRequestMessage)
	//fmt.Printf("context is : %#v",c)
	// fmt.Printf("context is : %v",c.Elements)
	// fmt.Printf("data : %v", cmd.Data)

	//fmt.Printf("cmd ctext is %v \n", c.Elements)

	interactor := usecases.NewContactAgentInteractor(memory.NewInMemoryHouseRepository())
	response := interactor.Handle(cmd.ContactAgentRequestMessage)
	fmt.Printf("the response is : %#v \n", response)

	return nil
}
