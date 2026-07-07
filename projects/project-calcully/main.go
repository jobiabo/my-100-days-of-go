package main

import (
	handler "calcully/handlers"
	"log"
	"net/http"
)

const powers = "㎡"

func main() {
	// fmt.Print("please enter numbers calculate: ")

	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')
	// input = strings.TrimRight(input, "\n")
	// fmt.Println(input)

	// switch {
	// case strings.HasPrefix(input, "sqrt"):
	// 	result, err := calcully.Sqrt(input)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(result)

	// case strings.HasPrefix(input, "sin"):
	// 	result, err := calcully.Sin(input)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(result)

	// case strings.HasPrefix(input, "cos"):
	// 	result, err := calcully.Cos(input)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(result)

	// case strings.HasPrefix(input, "tan"):
	// 	result, err := calcully.Tan(input)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(result)

	// case strings.HasPrefix(input, "log"):
	// 	result, err := calcully.Log(input)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(result)

	// case strings.Contains(input, powers):
	// 	result, err := (calcully.Power())
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(result)

	// default:
	// 	number, inString := (calcully.Check(input))

	// 	result, err := (calcully.Solve(inString, number))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	fmt.Println(result)
	// }
	// Serve everything in the template/static folder under /static/

	// mux := http.NewServeMux()
	// fs := http.FileServer(http.Dir("./template"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// mux.HandleFunc("/", handler.PageHandler)

	fs := http.FileServer(http.Dir("static")) // ✅ matches where IMG_9139.JPG actually is
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler.SplashPageHandler)
	http.HandleFunc("/home", handler.HomePageHandler)
	http.HandleFunc("/calculator", handler.PageHandler)
	http.HandleFunc("/userguide", handler.UserGuide)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/features", handler.FeatureHandlers)

	log.Println("server running at http://localhost:9000")
	// http.ListenAndServe(":9000", mux)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

}
