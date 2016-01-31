Why you should use go-boilerplate
-Robust 
-Simple design and flexible (help you manage your project in modules and you can also make use for all open source golang libraries in your project , it does not create anykind of hinderance for any library , which give you full control.)
-Modular and Replace-able modules (like for url routing it uses martini you can use gorilla-mux or anything else you want , similarly use orm , template engine of your choice or anything you can imagine )
- Dev and Live mode option to help you easily run and test it in development mode.
- Inbuilt different components to make you go with your application within minutes. see list of inbuilt components
- MVT based architecture . 



Framework Architecture
- MVT based architecure 
- Put all your modules in "apps" folder , like we have a "test" module with Models.go and Views.go
- "conf" folder contains the default config.ini file , you can use dev.ini file as well to define configurations in development environment.
- "static" contains all the static files you need for your project , access them directly with your website url and folder or file inside static . For Example : http://localhost:3000/test/1.jpg will display a file present at path static/test/1.jpg 
- "templates" 

Components 
-Manage configurations , the config module : Edit or create configuration variables easily , it make use of toml, and you can also add your own config variables as per your needs .   https://github.com/BurntSushi/toml
- Manage project dependencies : rely on godep to maintain dependencies https://github.com/tools/godep
- URL management and routing : for url management and routing it make use of martini  https://github.com/go-martini/martini
- ORM : use xorm , one of the best orm available for golang  https://github.com/go-xorm/xorm
- Templating : use mustache for templating , one of the fastest , easy to use and efficient templating engine 




