package config

import (
	"fmt"
	ini "gopkg.in/ini.v1"
	"log"
)

func Example() {
	cfg, err := ini.Load("./config.ini")
	checkErr(err)

	// get root key && value
	fmt.Println(cfg.Section("").Key("app_name").String())
	fmt.Println(cfg.Section("").Key("log_level").String())

	// get [mysql] content
	mysql := cfg.Section("mysql")
	port, _ := mysql.Key("port").Int()
	fmt.Println(mysql.Key("ip").String(), port, mysql.Key("user").String(), mysql.Key("password").String(), mysql.Key("database").String())

	// get all sections' name
	names := cfg.SectionStrings()
	fmt.Println(names)

	// get all keys
	vals := cfg.Section("").Keys()
	keys := cfg.Section("").KeyStrings()
	fmt.Printf("%v => %v\n", keys, vals)

	log_ip := cfg.Section("").Key("log_ip").Value()
	log_port := cfg.Section("").Key("log_port").Value()
	fmt.Printf("%v => %v\n", log_ip, log_port)

	// create new key && value, then save to file with indent format.
	cfg = ini.Empty()
	defaultSection := cfg.Section("mongodb")
	defaultSection.NewKey("ip", "127.0.0.1")
	defaultSection.NewKey("port", "45453")
	err = cfg.SaveToIndent("new.config.ini", "\t")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
