package main

import (
	"log"
	"net/http"
)

type welcomeHandler struct {
}

func (h *welcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Hey Coders!</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					text-align: center;
				}

				h1 {
					color: #333;
				}

				#cool-box {
					width: 200px;
					height: 200px;
					background-color: #00bcd4;
					margin: 20px auto;
					transition: transform 0.5s ease-in-out;
				}

				.cool-effect {
					transform: rotate(360deg);
				}
			</style>
		</head>
		<body>
			<h1>Welcome to Go!</h1>
			<div id="cool-box"></div>
			<button onclick="toggleCoolEffect()">Toggle Some Effect</button>
			<script>
				function toggleCoolEffect() {
					var box = document.getElementById('cool-box');
					box.classList.toggle('cool-effect');
				}
			</script>
		</body>
		</html>
	`))
}

func main() {
	http.Handle("/", &welcomeHandler{})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
