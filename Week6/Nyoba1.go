package main 
import "fmt" 
const pi float64 = 3.14 
 
func volume(r,t float64) float64{ 
    return pi * (r * r) * t 
} 
 
func massa(r,t,p float64) float64{ 
    var vol float64 = volume(r,t) 
    return vol * p 
} 
 
func display(m1, m2 float64) { 
    var selisih float64 
    if m1 == m2 { 
        fmt.Println("BALANCE") 
    }else{ 
        selisih = m1 - m2 
        if selisih < 0 { 
            selisih = - selisih 
        } 
        fmt.Println(selisih) 
    } 
} 

func main(){ 
    var r float64                         
    var tKiri, tKanan float64            
    var mjKiri, mjKanan float64         
    var massaKiri, massaKanan float64    

    fmt.Scan(&r) 
    fmt.Scan(&tKiri, &mjKiri) 
    fmt.Scan(&tKanan ,&mjKanan) 
    massaKiri = massa(r,tKiri,mjKiri) 
    massaKanan = massa(r,tKanan,mjKanan) 
    display(massaKiri,massaKanan) 
}