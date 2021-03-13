package giveDuck

import (
	"fmt"
    "strings"
	"github.com/jonas747/dcmd"
	"github.com/jonas747/yagpdb/commands"
)

var Command = &commands.YAGCommand{
	Cooldown:    5,
	CmdCategory: commands.CategoryFun,
	Name:        "GiveDuck",
	Description: "Generates a Duck Image using a random duck API",
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
	
	link := "https://random-d.uk/api/"
	c := randInt 10
	if lt c 7 }}
		link = joinStr "" link (toString (randInt 1 130) ) ".jpg"
	else
		link = joinStr "" link (toString (randInt 1 29)) ".gif"
	end
	returndata := (cembed "image" (sdict "url" $link))
	return $returndata, nil
	}
}

