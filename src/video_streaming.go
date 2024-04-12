package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello World")
// }

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/pion/webrtc/v3"
// 	"github.com/rs/cors"
// )

// func main() {
// 	// Inicializar el router y el middleware CORS
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", indexHandler)
// 	mux.HandleFunc("/offer", offerHandler)
// 	handler := cors.Default().Handler(mux)

// 	// Iniciar el servidor HTTP
// 	log.Println("Server started at :5000")
// 	if err := http.ListenAndServe(":5000", handler); err != nil {
// 		log.Fatalf("Failed to start server: %+v", err)
// 	}
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	// Sirve la página HTML aquí
// 	w.Header().Set("Content-Type", "text/html")
// 	http.ServeFile(w, r, "index.html")
// }

// func offerHandler(w http.ResponseWriter, r *http.Request) {
// 	// Aquí se manejaría la oferta SDP enviada por el cliente
// 	// Por simplificación, este código no está incluido
// 	// Deberías parsear la oferta SDP, crear una respuesta, y configurar los tracks

// 	// Configuración básica de WebRTC
// 	m := webrtc.MediaEngine{}
// 	if err := m.RegisterDefaultCodecs(); err != nil {
// 		log.Fatalf("Failed to register default codecs: %+v", err)
// 	}

// 	api := webrtc.NewAPI(webrtc.WithMediaEngine(&m))
// 	pc, err := api.NewPeerConnection(webrtc.Configuration{})
// 	if err != nil {
// 		log.Fatalf("Failed to create new PeerConnection: %+v", err)
// 	}

// 	// Aquí añadirías tracks al PeerConnection, por ejemplo, para video
// 	// track, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: "video/VP8"}, "video", "pion")
// 	// if err != nil {
// 	//     log.Fatalf("Failed to create video track: %+v", err)
// 	// }
// 	// _, err = pc.AddTrack(track)
// 	// if err != nil {
// 	//     log.Fatalf("Failed to add track to PeerConnection: %+v", err)
// 	// }

// 	// Código para manejar SDP Offer/Answer aquí

// 	// Devolver la respuesta SDP al cliente
// 	w.Header().Set("Content-Type", "application/json")
// 	// json.NewEncoder(w).Encode(answer)
// }
