package controller

import (
	"fmt"
	"time"
	"strings"
	lg_tv "github.com/0187773933/LGTVController/v1/controller"
	tg_tv_types "github.com/0187773933/LGTVController/v1/types"
	vizio_tv "github.com/0187773933/VizioController/v1/controller"
	hdmi_cec "github.com/0187773933/HDMICEC/v1/controller"
	ir "github.com/0187773933/IRController/v1/controller"
	types "github.com/0187773933/TVController/v1/types"
	utils "github.com/0187773933/TVController/v1/utils"
	try "github.com/manucorporat/try"
)

type Controller struct {
	Type string `yaml:"type"`
	Config *types.ConfigFile `yaml:"-"`
	LG *lg_tv.Controller `yaml:"-"`
	VIZIO *vizio_tv.Controller `yaml:"-"`
	HDMICEC *hdmi_cec.Controller `yaml:"-"`
	IR *ir.Controller `yaml:"-"`
	LG_Ready bool `yaml:"-"`
	VIZIO_Ready bool `yaml:"-"`
	SAMSUNG_Ready bool `yaml:"-"`
	IR_Ready bool `yaml:"-"`
	HDMI_CEC_Ready bool `yaml:"-"`
}

func New( config *types.ConfigFile ) ( result *Controller ) {
	result = &Controller{ Config: config }
	brand_lower := strings.ToLower( config.Brand )
	result.Type = brand_lower
	switch result.Type{
		case "lg":
			try.This( func() {
				lg_tv_config := &tg_tv_types.ConfigFile{
					TVIP: config.IP ,
					TVMAC: config.MAC ,
					WebSocketPort: config.LGWebSocketPort ,
					ClientKey: config.LGClientKey ,
					TimeoutSeconds: config.TimeoutSeconds ,
				}
				result.LG = lg_tv.New( lg_tv_config )
				result.LG_Ready = true
			}).Catch(func(e try.E) {
				fmt.Println( e )
				fmt.Println( "Failed to add lg" )
				result.LG_Ready = false
			})
			break;
		case "samsung":
			try.This( func() {
				fmt.Println( "samsung === todo" )
				result.SAMSUNG_Ready = true
			}).Catch(func(e try.E) {
				fmt.Println( e )
				fmt.Println( "Failed to add samsung" )
				result.SAMSUNG_Ready = false
			})
			break;
		case "vizio":
			try.This( func() {
				result.VIZIO = vizio_tv.New( config.IP , config.VizioAuthToken )
				result.VIZIO_Ready = true
			}).Catch(func(e try.E) {
				fmt.Println( e )
				fmt.Println( "Failed to add vizio" )
				result.VIZIO_Ready = false
			})
			break;
		case "hdmicec":
			try.This( func() {
				hcec := hdmi_cec.New()
				result.HDMICEC = &hcec
				result.HDMI_CEC_Ready = true
			}).Catch(func(e try.E) {
				fmt.Println( e )
				fmt.Println( "Failed to add hdmi cec adapter" )
				result.HDMI_CEC_Ready = false
			})
			break;
		case "ir":
			try.This( func() {
				// ir_config := utils.ParseIRConfig( config.IRConfig )
				x := ir.New( &config.IRConfig )
				result.IR = &x
				result.IR_Ready = true
			}).Catch( func( e try.E ) {
				fmt.Println( e )
				fmt.Println( "Failed to add ir adapter" )
				result.IR_Ready = false
			})
			break;
		case "ir+hdmicec":
			try.This( func() {
				hcec := hdmi_cec.New()
				result.HDMICEC = &hcec
				result.HDMI_CEC_Ready = true
			}).Catch(func(e try.E) {
				fmt.Println( e )
				fmt.Println( "Failed to add hdmi cec adapter" )
				result.HDMI_CEC_Ready = false
			})
			try.This( func() {
				// ir_config := utils.ParseIRConfig( config.IRConfigPath )
				x := ir.New( &config.IRConfig )
				result.IR = &x
				result.IR_Ready = true
			}).Catch( func( e try.E ) {
				fmt.Println( e )
				fmt.Println( "Failed to add ir adapter" )
				result.IR_Ready = false
			})
			break;
	}
	return
}

func ( c *Controller ) WakeOnLAN() {
	utils.WakeOnLan( c.Config.MAC )
}

func ( c *Controller ) Reset() {
	switch c.Type{
		case "lg":
			fmt.Println( "lg === todo" )
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			fmt.Println( "vizio === todo" )
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === Reset()" )
			power_status := c.HDMICEC.GetPowerStatus()
			fmt.Println( "Already On ===" , power_status )
			time.Sleep( 1200 * time.Millisecond )
			if power_status == false {
				fmt.Println( "Powering On" )
				c.HDMICEC.PowerOn()
			}
			fmt.Printf( "Setting HDMI %d\n" , c.Config.DefaultInput )
			c.HDMICEC.SelectHDMI( c.Config.DefaultInput )
			if c.IR_Ready {
				exit_code := c.Config.IRConfig.Remotes[ c.Config.IRConfig.DefaultRemote ].Keys[ "exit" ].Code
				if exit_code != "" {
					fmt.Println( "Sending IR Code" , exit_code )
					c.IR.Transmit( exit_code )
					c.IR.Transmit( exit_code )
				}
				volume_down := c.Config.IRConfig.Remotes[ c.Config.IRConfig.DefaultRemote ].Keys[ "volume_down" ].Code
				volume_up := c.Config.IRConfig.Remotes[ c.Config.IRConfig.DefaultRemote ].Keys[ "volume_up" ].Code
				if volume_down != "" && volume_up != "" {
					fmt.Println( "Resetting Volume" )
					for i := 0; i < c.Config.VolumeResetLimit; i++ {
						fmt.Printf( "Volume Down %d of %d\n" , i , c.Config.VolumeResetLimit )
						c.IR.Transmit( volume_down )
						time.Sleep( 500 * time.Millisecond )
					}
					fmt.Println( "Done With Volume Down" )
					time.Sleep( 500 * time.Millisecond )
					for i := 0; i < c.Config.DefaultVolume; i++ {
						fmt.Printf( "Volume Up %d of %d\n" , i , c.Config.DefaultVolume )
						c.IR.Transmit( volume_up )
						time.Sleep( 500 * time.Millisecond )
					}
				}
			}
			break;
	}
	return
}

func ( c *Controller ) Prepare() {
	fmt.Println( "Prepare()" )
	switch c.Type{
		case "lg":
			c.PowerOn()
			c.SetInput( c.Config.DefaultInput )
			c.SetVolume( c.Config.DefaultVolume )
			break;
		case "samsung":
			c.PowerOn()
			c.SetInput( c.Config.DefaultInput )
			c.SetVolume( c.Config.DefaultVolume )
			break;
		case "vizio":
			c.PowerOn()
			c.SetInput( c.Config.DefaultInput )
			c.SetVolume( c.Config.DefaultVolume )
			break;
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			c.PowerOn()
			break;
		case "hdmicec":
		case "ir+hdmicec":
			status := c.Status()
			utils.PrettyPrint( status )
			if status.Power == false {
				fmt.Println( "Prepare() --> Power == false" )
				time.Sleep( 1200 * time.Millisecond )
				c.PowerOn()
				time.Sleep( 1200 * time.Millisecond )
				c.SetInput( c.Config.DefaultInput )
				time.Sleep( 1200 * time.Millisecond )
				c.SetVolume( c.Config.DefaultVolume )
				return
			}
			if status.HDMIInput != c.Config.DefaultInput {
				fmt.Println( "Prepare() --> Resetting HDMI Input" )
				c.SetInput( c.Config.DefaultInput )
			}
			if status.Volume != -1 && status.Volume != c.Config.DefaultVolume {
				fmt.Println( "Prepare() --> Resetting Volume" )
				c.SetVolume( c.Config.DefaultVolume )
			}
			fmt.Println( "Prepare() --> Done" )
			break;
	}
	return
}

func ( c *Controller ) ResetVideo() {
	switch c.Type{
		case "lg":
			fmt.Println( "lg === todo" )
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			fmt.Println( "vizio === todo" )
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === ResetVideo()" )
			power_status := c.HDMICEC.GetPowerStatus()
			fmt.Println( "Already On ===" , power_status )
			time.Sleep( 1200 * time.Millisecond )
			if power_status == false {
				fmt.Println( "Powering On" )
				c.HDMICEC.PowerOn()
			}
			fmt.Printf( "Setting HDMI %d\n" , c.Config.DefaultInput )
			c.HDMICEC.SelectHDMI( c.Config.DefaultInput )
			break;
	}
	return
}

func ( c *Controller ) QuickResetVideo() {
	switch c.Type{
		case "lg":
			fmt.Println( "lg === todo" )
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			fmt.Println( "vizio === todo" )
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === QuickResetVideo()" )
			fmt.Println( "Powering On" )
			c.HDMICEC.PowerOn()
			fmt.Printf( "Setting HDMI %d\n" , c.Config.DefaultInput )
			c.HDMICEC.SelectHDMI( c.Config.DefaultInput )
			break;
	}
	return
}

type Status struct {
	Volume int `json:"volume"`
	Power bool `json:"power"`
	HDMIInput int `json:"hdmi_input"`
	HDMIVendor string `json:"hdmi_vendor"`
	HDMIOSDString string `json:"hdmi_osd_string"`
	HDMIPower bool `json:"hdmi_power"`
}

func ( c *Controller ) Status() ( result Status ) {
	switch c.Type{
		case "lg":
			volume_string := c.LG.API( "get_volume" )
			result.Volume = utils.StringToInt( volume_string )
			inputs_string := c.LG.API( "get_inputs" )
			fmt.Println( "lg === to do , unknown what get_inputs list is" , inputs_string )
			result.HDMIInput = -1
			// so you just try to access some endpoint , if its there , then tv is on
			read_result := c.LG.API( "get_volume" )
			if read_result == "error reading message" {
				result.Power = false
			} else if read_result == "timeout while reading message" {
				result.Power = false
			} else {
				result.Power = true
			}
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			result.Volume = c.VIZIO.VolumeGet()
			current_input := c.VIZIO.InputGetCurrent()
			switch current_input.Name {
				case "hdmi1":
					result.HDMIInput = 1
					break;
				case "hdmi2":
					result.HDMIInput = 2
					break;
				case "hdmi3":
					result.HDMIInput = 3
					break;
				case "hdmi4":
					result.HDMIInput = 4
					break;
				default:
					result.HDMIInput = -1
					break
			}
			x := c.VIZIO.PowerGetState()
			fmt.Println( "vizio === todo" , x )
			result.Power = false
			break;
		case "hdmicec":
			c.HDMICEC.PowerOn()
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			result.Volume = -1
			sources := c.HDMICEC.GetSources()
			for _ , source := range sources {
				if source.DeviceName == "TV" {
					result.Power = source.PowerStatus
				}
				if source.ActiveSource == true {
					// utils.PrettyPrint( source )
					result.HDMIInput = source.HDMIInput
					result.HDMIVendor = source.Vendor
					result.HDMIOSDString = source.OSDString
					result.HDMIPower = source.PowerStatus
				}
			}
			break;
	}
	return
}

func ( c *Controller ) PowerOn() {
	switch c.Type{
		case "lg":
			c.LG.API( "power_on" )
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			c.VIZIO.PowerOn()
			break;
		case "hdmicec":
			c.HDMICEC.PowerOn()
			break;
		case "ir":
			power := c.Config.IRConfig.Remotes[ c.Config.IRConfig.DefaultRemote ].Keys[ "power" ].Code
			c.IR.Transmit( power )
			break;
		case "ir+hdmicec":
			c.HDMICEC.PowerOn()
			break;
	}
	return
}

func ( c *Controller ) PowerOff() {
	switch c.Type{
		case "lg":
			c.LG.API( "power_off" )
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			c.VIZIO.PowerOff()
			break;
		case "hdmicec":
			c.HDMICEC.PowerOff()
			break;
		case "ir":
			power := c.Config.IRConfig.Remotes[ c.Config.IRConfig.DefaultRemote ].Keys[ "power" ].Code
			c.IR.Transmit( power )
			break;
		case "ir+hdmicec":
			c.HDMICEC.PowerOff()
			break;
	}
	return
}

func ( c *Controller ) GetPowerStatus() ( result bool ) {
	result = false
	switch c.Type{
		case "lg":
			// so you just try to access some endpoint , if its there , then tv is on
			read_result := c.LG.API( "get_volume" )
			if read_result == "error reading message" {
				result = false
			} else if read_result == "timeout while reading message" {
				result = false
			} else {
				result = true
			}
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			x := c.VIZIO.PowerGetState()
			fmt.Println( "vizio === todo" , x )
			break;
		case "hdmicec":
			result = c.HDMICEC.GetPowerStatus()
			break;
		case "ir":
			fmt.Println( "ir === impossible" )
			break;
		case "ir+hdmicec":
			result = c.HDMICEC.GetPowerStatus()
			break;
	}
	return
}

func ( c *Controller ) GetInput() ( result int ) {
	switch c.Type{
		case "lg":
			result_string := c.LG.API( "get_inputs" )
			fmt.Println( "lg === to do , unknown what get_inputs list is" , result_string )
			result = -1
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			current_input := c.VIZIO.InputGetCurrent()
			switch current_input.Name {
				case "hdmi1":
					result = 1
					break;
				case "hdmi2":
					result = 2
					break;
				case "hdmi3":
					result = 3
					break;
				case "hdmi4":
					result = 4
					break;
				default:
					result = -1
					break
			}
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === todo" )
			break;
	}
	return
}

func ( c *Controller ) SetInput( hdmi_input int ) {
	switch c.Type{
		case "lg":
			c.LG.API( "set_input" , tg_tv_types.Payload{
				"inputId": fmt.Sprintf( "HDMI-%d" , hdmi_input ) ,
			})
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			target_input := fmt.Sprintf( "HDMI-%d" , hdmi_input )
			fmt.Println( "setting vizio to :" , target_input )
			c.VIZIO.InputSet( target_input )
			break;
		case "hdmicec":
			fmt.Println( "hdmicec to :" , hdmi_input )
			c.HDMICEC.SelectHDMI( hdmi_input )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "hdmicec to :" , hdmi_input )
			c.HDMICEC.SelectHDMI( hdmi_input )
			break;
	}
	return
}

func ( c *Controller ) MuteOn() {
	switch c.Type{
		case "lg":
			c.LG.API( "set_mute" , tg_tv_types.Payload{
				"mute": true ,
			})
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			c.VIZIO.MuteOn()
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === todo" )
			break;
	}
	return
}

func ( c *Controller ) MuteOff() {
	switch c.Type{
		case "lg":
			c.LG.API( "set_mute" , tg_tv_types.Payload{
				"mute": false ,
			})
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			c.VIZIO.MuteOff()
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === todo" )
			break;
	}
	return
}

func ( c *Controller ) GetMute() ( result bool ) {
	switch c.Type{
		case "lg":
			// audio_status := tv.LG.API( "get_audio_status" )
			fmt.Println( "GetMute() lg === to do" )
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			x := c.VIZIO.AudioGetSetting( "mute" )
			fmt.Println( "vizio === todo" , x )
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === todo" )
			break;
	}
	return
}

func ( c *Controller ) GetVolume() ( result int ) {
	switch c.Type{
		case "lg":
			result_string := c.LG.API( "get_volume" )
			result = utils.StringToInt( result_string )
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			result = c.VIZIO.VolumeGet()
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === todo" )
			break;
	}
	return
}

func ( c *Controller ) SetVolume( volume_level int ) {
	switch c.Type{
		case "lg":
			c.LG.API( "set_input" , tg_tv_types.Payload{
				"volume": volume_level ,
			})
			break;
		case "samsung":
			fmt.Println( "samsung === todo" )
			break;
		case "vizio":
			c.VIZIO.VolumeSet( volume_level )
			break;
		case "hdmicec":
			fmt.Println( "hdmicec === todo" )
			break;
		case "ir":
			fmt.Println( "ir === todo" )
			break;
		case "ir+hdmicec":
			fmt.Println( "ir+hdmicec === SetVolume()" )
			if c.IR_Ready {
				volume_down := c.Config.IRConfig.Remotes[ c.Config.IRConfig.DefaultRemote ].Keys[ "volume_down" ].Code
				volume_up := c.Config.IRConfig.Remotes[ c.Config.IRConfig.DefaultRemote ].Keys[ "volume_up" ].Code
				if volume_down != "" && volume_up != "" {
					fmt.Println( "Resetting Volume" )
					for i := 0; i < c.Config.VolumeResetLimit; i++ {
						fmt.Printf( "Volume Down %d of %d\n" , i , c.Config.VolumeResetLimit )
						c.IR.Transmit( volume_down )
						time.Sleep( 500 * time.Millisecond )
					}
					fmt.Println( "Done With Volume Down" )
					time.Sleep( 500 * time.Millisecond )
					for i := 0; i < c.Config.DefaultVolume; i++ {
						fmt.Printf( "Volume Up %d of %d\n" , i , c.Config.DefaultVolume )
						c.IR.Transmit( volume_up )
						time.Sleep( 500 * time.Millisecond )
					}
				}
			}
			break;
	}
	return
}