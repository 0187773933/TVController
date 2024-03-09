package main

import (
	"fmt"
	"os"
	"path/filepath"
	utils "github.com/0187773933/TVController/v1/utils"
	tv_controller "github.com/0187773933/TVController/v1/controller"
)

func main() {
	config_file_path , _ := filepath.Abs( "./config.yaml" )
	if len( os.Args ) > 1 { config_file_path , _ = filepath.Abs( os.Args[ 1 ] ) }
	config := utils.ParseConfig( config_file_path )
	fmt.Printf( "Loaded Config File From : %s\n" , config_file_path )
	utils.PrettyPrint( config )
	tv := tv_controller.New( &config )
	fmt.Println( tv )
	tv.Prepare()
	// status := tv.Status()
	// utils.PrettyPrint( status )
}
