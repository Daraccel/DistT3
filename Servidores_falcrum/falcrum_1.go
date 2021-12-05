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

var conn = ":50052"

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
	return &pb.ServidorInformante{
		X: 1,
		Y: 1,
		Z: 1,
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