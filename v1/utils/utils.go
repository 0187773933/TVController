package utils

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"net"
	binary "encoding/binary"
	json "encoding/json"
	ioutil "io/ioutil"
	yaml "gopkg.in/yaml.v2"
	types "github.com/0187773933/TVController/v1/types"
	ir_types "github.com/0187773933/IRController/v1/types"
)

func IToB( v uint64 ) []byte {
	b := make( []byte , 8 )
	binary.BigEndian.PutUint64( b , v )
	return b
}

func StringToInt( input string ) ( result int ) {
	result , _ = strconv.Atoi( input )
	return
}

func GetFormattedTimeString() ( result string ) {
	location , _ := time.LoadLocation( "America/New_York" )
	time_object := time.Now().In( location )
	month_name := strings.ToUpper( time_object.Format( "Jan" ) )
	milliseconds := time_object.Format( ".000" )
	date_part := fmt.Sprintf( "%02d%s%d" , time_object.Day() , month_name , time_object.Year() )
	time_part := fmt.Sprintf( "%02d:%02d:%02d%s" , time_object.Hour() , time_object.Minute() , time_object.Second() , milliseconds )
	result = fmt.Sprintf( "%s === %s" , date_part , time_part )
	return
}

func GetFormattedTimeStringOBJ() ( result_string string , result_time time.Time ) {
	location , _ := time.LoadLocation( "America/New_York" )
	result_time = time.Now().In( location )
	month_name := strings.ToUpper( result_time.Format( "Jan" ) )
	milliseconds := result_time.Format( ".000" )
	date_part := fmt.Sprintf( "%02d%s%d" , result_time.Day() , month_name , result_time.Year() )
	time_part := fmt.Sprintf( "%02d:%02d:%02d%s" , result_time.Hour() , result_time.Minute() , result_time.Second() , milliseconds )
	result_string = fmt.Sprintf( "%s === %s" , date_part , time_part )
	return
}

func FormatTime( input_time *time.Time ) ( result string ) {
	location , _ := time.LoadLocation( "America/New_York" )
	time_object := input_time.In( location )
	month_name := strings.ToUpper( time_object.Format( "Jan" ) )
	milliseconds := time_object.Format( ".000" )
	date_part := fmt.Sprintf( "%02d%s%d" , time_object.Day() , month_name , time_object.Year() )
	time_part := fmt.Sprintf( "%02d:%02d:%02d%s" , time_object.Hour() , time_object.Minute() , time_object.Second() , milliseconds )
	result = fmt.Sprintf( "%s === %s" , date_part , time_part )
	return
}

func WakeOnLan( mac_address string ) {
	mac_bytes , _ := net.ParseMAC( mac_address )
	magic_packet := []byte{}
	for i := 0; i < 6; i++ {
		magic_packet = append( magic_packet , 0xFF )
	}
	for i := 0; i < 16; i++ {
		magic_packet = append( magic_packet , mac_bytes... )
	}
	addr := &net.UDPAddr{
		IP: net.IPv4bcast ,
		Port: 9 ,
	}
	conn , _ := net.DialUDP( "udp" , nil , addr )
	defer conn.Close()
	conn.Write( magic_packet )
}

func PrettyPrint( input interface{} ) {
	jd , _ := json.MarshalIndent( input , "" , "  " )
	fmt.Println( string( jd ) )
}

func WriteJSON( filePath string , data interface{} ) {
	file, _ := json.MarshalIndent( data , "" , " " )
	_ = ioutil.WriteFile( filePath , file , 0644 )
}

func ParseConfig( file_path string ) ( result types.ConfigFile ) {
	file , _ := ioutil.ReadFile( file_path )
	error := yaml.Unmarshal( file , &result )
	if error != nil { panic( error ) }
	return
}

func ParseIRConfig( file_path string ) ( result ir_types.ConfigFile ) {
	file , _ := ioutil.ReadFile( file_path )
	error := yaml.Unmarshal( file , &result )
	if error != nil { panic( error ) }
	return
}