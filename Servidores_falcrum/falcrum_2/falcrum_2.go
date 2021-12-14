package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	// "log"
	// "math/rand"
	"net"
	// "sync"
	// "time"

	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
)

// coord de servidor
var conn = ":50053"

// var para vector coord a sumar en planeta-lugar modificado
var c_x int32 = 0
var c_y int32 = 1
var c_z int32 = 0

type coords struct {
	coor_x, coor_y, coor_z int32
}

type infoplanet struct {
	rebels              int32
	cordx, cordy, cordz int32
	server              string
}

var mapa_citinfo map[string]infoplanet

var mapeo_plan_coord map[string]coords

type server struct {
	pb.UnimplementedFuncionesServiceServer
}

func append_line_archivo(nombre string, nueva_linea string) {
	file, err := os.OpenFile(nombre, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("err")
		return
	}

	_, err = fmt.Fprintln(file, nueva_linea)
	if err != nil {
		fmt.Println("err")
		file.Close()
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file append successesfully")
}

//cada vez que se corre esto, se crea el archivo vacio
func crear_archivo(nombre string) {
	file, err := os.Create(nombre)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
		fmt.Println("File creada")
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("File has been created successfully!.")
}

func Up_num(linea_nueva string, buscar string, nombre string) {
	input, err := ioutil.ReadFile(nombre)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, buscar) {
			lines[i] = linea_nueva
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(nombre, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//                   ciudad        planeta
func Up_numLei(buscar string, nombre string) (string, int32) {
	input, err := ioutil.ReadFile(nombre + ".txt")
	if err != nil {
		fmt.Println(err)
		return "error", 0
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, buscar) {
			este := strings.Split(lines[i], " ")
			x, _ := strconv.Atoi(este[2])
			return este[1], int32(x)
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(nombre, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
		return "error", 0
	}
	return "error", 0
}

func camb_nomb(val string, buscar string, nombre string) {
	input, err := ioutil.ReadFile(nombre)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, buscar) {
			fmt.Println(line)
			lin := strings.Split(string(lines[i]), " ")

			for x, y := range lin {
				if strings.Contains(y, buscar) {
					lin[x] = val
				}
			}
			linea_nueva := strings.Join(lin, " ")
			lines[i] = linea_nueva
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(nombre, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *server) Inf_BroServCoord(ctx context.Context, req *pb.BrokerServidorCoord) (*pb.ServidorBrokerCoord, error) {
	var nuevo_x, nuevo_y, nuevo_z int32

	coordenadas, ok := mapeo_plan_coord[req.Planeta]

	if ok {
		fmt.Println("Ok: TRUE")
		nuevo_x = coordenadas.coor_x
		nuevo_y = coordenadas.coor_y
		nuevo_z = coordenadas.coor_y
	} else {
		fmt.Println("creano map")
		nuevo_x = 0
		nuevo_y = 0
		nuevo_z = 0
	}

	return &pb.ServidorBrokerCoord{
		X: nuevo_x,
		Y: nuevo_y,
		Z: nuevo_z,
	}, nil
}

func (c *server) InfBroful(ctx context.Context, req *pb.BrokerFulcrum) (*pb.FulcrumBroker, error) {
	var rebels, cordex, cordey, cordez int32

	info, ok := mapeo_plan_coord[req.Nombreplaneta]

	if ok {
		fmt.Println("Buscando ciudades del planeta")
		city, rebs := Up_numLei(req.Nombreciudad, req.Nombreplaneta)
		if city != "error" {
			rebels = rebs
			cordex = info.coor_x
			cordey = info.coor_y
			cordez = info.coor_z
		} else {
			fmt.Println("Si no esta en nuestros archivos es porque no existe")

		}

	}
	return &pb.FulcrumBroker{
		X:            cordex,
		Y:            cordey,
		Z:            cordez,
		Rebeldes:     rebels,
		Nombreserver: conn,
	}, nil

}

func (s *server) InfServ(ctx context.Context, req *pb.InformanteServidor) (*pb.ServidorInformante, error) {
	accion := req.Accion
	planeta := req.PlanetaAfectado
	ciudad := req.CiudadAfectada
	val := req.NuevoValor

	nom := planeta + ".txt"
	nom2 := "registro_" + planeta + ".txt"
	str := planeta + " " + ciudad + " " + val
	str2 := accion + " " + planeta + " " + ciudad + " " + val

	fmt.Println("Nom: " + nom)
	fmt.Println("Accion: " + accion)
	fmt.Println("planeta: " + planeta)
	fmt.Println("ciudad: " + ciudad)
	fmt.Println("val: " + val)

	///////////-----Seccion de archivos de planetas-----///////////
	if accion == "AddCity" {
		//chequear si existe archivo y anidir linea

		if _, err := os.Stat(nom); err == nil {
			fmt.Println("Archivo existe, creando...")

			append_line_archivo(nom, str)
		} else if err != nil {
			fmt.Println("No se encontro archivo, creando...")
			crear_archivo(nom)
			append_line_archivo(nom, str)
		} else {
			fmt.Println("Error en el archivo a escribir")
		}

	} else if accion == "UpdateName" {
		//chequear si existe archivo y anidir linea
		if _, err := os.Stat(nom); err == nil {
			fmt.Println("Archivo existe")

			camb_nomb(val, ciudad, nom)
		} else if err != nil {
			fmt.Println("No se encontro archivo, creando...")
			crear_archivo(nom)
			camb_nomb(val, ciudad, nom)
		} else {
			fmt.Println("Error en el archivo a escribir")
		}

	} else if accion == "UpdateNumber" {
		//chequear si existe archivo y anidir linea
		str := planeta + " " + ciudad + " " + val
		if _, err := os.Stat(nom); err == nil {
			fmt.Println("Archivo existe")

			Up_num(str, ciudad, nom)
		} else if err != nil {
			fmt.Println("No se encontro archivo, creando...")
			crear_archivo(nom)
			Up_num(str, ciudad, nom)
		} else {
			fmt.Println("Error en el archivo a escribir")
		}

	} else if accion == "DeleteCity" {
		//chequear si existe archivo y anidir linea
		if _, err := os.Stat(nom); err == nil {
			fmt.Println("Archivo existe")

			Up_num("", ciudad, nom)
		} else if err != nil {
			fmt.Println("No se encontro archivo, creando...")
			crear_archivo(nom)
			Up_num("", ciudad, nom)
		} else {
			fmt.Println("Error en el archivo a escribir")
		}

	}
	///////////-----------------------------------------///////////

	///////////-----Seccion de registros de planetas----///////////
	if _, err := os.Stat(nom2); err == nil {
		fmt.Println("Archivo existe, creando...")

		append_line_archivo(nom2, str2)
	} else if err != nil {
		fmt.Println("No se encontro archivo, creando...")
		crear_archivo(nom2)
		append_line_archivo(nom2, str2)
	} else {
		fmt.Println("Error en el archivo a escribir")
	}
	///////////-----------------------------------------///////////

	// agregar a map con key nombre planeta

	coordenadas, ok := mapeo_plan_coord[planeta]

	if ok {
		fmt.Println("Ok: TRUE")
		fmt.Println(coordenadas.coor_x)
		fmt.Println(coordenadas.coor_y)
		fmt.Println(coordenadas.coor_z)
		nuevo_x := coordenadas.coor_x + c_x
		nuevo_y := coordenadas.coor_y + c_y
		nuevo_z := coordenadas.coor_y + c_z
		mapeo_plan_coord[planeta] = coords{nuevo_x, nuevo_y, nuevo_z}
	} else {
		fmt.Println("creano map")
		fmt.Println(c_x)
		fmt.Println(c_y)
		fmt.Println(c_z)

		mapeo_plan_coord[planeta] = coords{c_x, c_y, c_z}
	}
	// fmt.Println()

	return &pb.ServidorInformante{
		X: mapeo_plan_coord[planeta].coor_x,
		Y: mapeo_plan_coord[planeta].coor_y,
		Z: mapeo_plan_coord[planeta].coor_z,
	}, nil
}

func main() {
	mapeo_plan_coord = make(map[string]coords)

	listner, err := net.Listen("tcp", conn)

	if err != nil {
		panic("No se pudo crear coneccion tcp de servidor serv 1 " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterFuncionesServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("No se pudo inicializar servidor serv 1 " + err.Error())
	}

}
