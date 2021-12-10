package main 


import (
	"io/ioutil"
	"strings"
	"context"
	"fmt"
	"os"
	// "log"
	// "math/rand"
	"net"
	// "sync"
	// "time"

	pb "github.com/Daraccel/DistT3/proto"
	"google.golang.org/grpc"
)

// coord de servidor
var conn 	= ":50052"

// var para vector coord a sumar en planeta-lugar modificado 
var c_x int32	= 1
var c_y int32	= 0
var c_z	int32	= 0

type coords struct {
	coor_x,coor_y,coor_z int32
}

var mapeo_plan_coord map[string]coords

type server struct{
	pb.UnimplementedFuncionesServiceServer
}

func append_line_archivo(nombre string, nueva_linea string) {
	file,err := os.OpenFile(nombre, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("err")
		return
	}

	_, err = fmt.Fprintln(file,nueva_linea)
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

func crear_archivo(nombre string) {
	//cada vez que se corre esto, se crea el archivo vacio
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

	output := strings.Join(lines,"\n")
	err = ioutil.WriteFile(nombre, []byte(output), 0644)
	if err!= nil {
		fmt.Println(err)
		return
	}
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

	output := strings.Join(lines,"\n")
	err = ioutil.WriteFile(nombre, []byte(output), 0644)
	if err!= nil {
		fmt.Println(err)
		return
	}
}


// Por modificar esto, x,y,z es por archivo y aumenta con cada uso
func (s *server) InfServ(ctx context.Context, req *pb.InformanteServidor) (*pb.ServidorInformante, error){
	accion 	:= req.Accion
	planeta := req.PlanetaAfectado
	ciudad	:= req.CiudadAfectada
	val		:= req.NuevoValor

	nom := planeta + ".txt"
	

	if accion == "AddCity"{
		//chequear si existe archivo y anidir linea
		str := planeta + " " + ciudad + " " + val
		if _, err := os.Stat(nom); err == nil {
			fmt.Println("Archivo existe, creando...")
			
			append_line_archivo(nom, str)
		} else if err !=  nil{
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
			
			camb_nomb(val,ciudad, nom)
		} else if err !=  nil{
			fmt.Println("No se encontro archivo, creando...")
			crear_archivo(nom)
			camb_nomb(val,ciudad, nom)
		} else {
			fmt.Println("Error en el archivo a escribir")
		}

	} else if accion == "UpdateNumber"{
		//chequear si existe archivo y anidir linea
		str := planeta + " " + ciudad + " " + val
		if _, err := os.Stat(nom); err == nil {
			fmt.Println("Archivo existe")
			
			Up_num(str,ciudad, nom)
		} else if err !=  nil{
			fmt.Println("No se encontro archivo, creando...")
			crear_archivo(nom)
			Up_num(str,ciudad, nom)
		} else {
			fmt.Println("Error en el archivo a escribir")
		}

	} else if accion == "DeleteCity" {
		//chequear si existe archivo y anidir linea
		if _, err := os.Stat(nom); err == nil {
			fmt.Println("Archivo existe")
			
			Up_num("",ciudad, nom)
		} else if err !=  nil{
			fmt.Println("No se encontro archivo, creando...")
			crear_archivo(nom)
			Up_num("",ciudad, nom)
		} else {
			fmt.Println("Error en el archivo a escribir")
		}

	}

	// agregar a map con key nombre planeta
	mapeo_plan_coord = make(map[string]coords)

	coordenadas, ok := mapeo_plan_coord[planeta]

	if ok {
		nuevo_x := coordenadas.coor_x + c_x
		nuevo_y := coordenadas.coor_y + c_y
		nuevo_z := coordenadas.coor_y + c_z
		mapeo_plan_coord[planeta] = coords{nuevo_x,nuevo_y,nuevo_z}
	} else{
		mapeo_plan_coord[planeta] = coords{c_x,c_y,c_z}
	}
	fmt.Println()

	return &pb.ServidorInformante{
		X: mapeo_plan_coord[planeta].coor_x,
		Y: mapeo_plan_coord[planeta].coor_y,
		Z: mapeo_plan_coord[planeta].coor_z,
	}, nil
}

func main() {
	listner, err := net.Listen("tcp",conn)
	
	if err != nil {
		panic ("No se pudo crear coneccion tcp de servidor serv 1 " + err.Error())
	}
	
	serv := grpc.NewServer() 
	pb.RegisterFuncionesServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("No se pudo inicializar servidor serv 1 " + err.Error())
	}

}