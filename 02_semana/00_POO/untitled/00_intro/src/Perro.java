public class Perro {
    private Boolean estaEnAdopcioon;
    private  String raza;
    private Integer anioNacimiento;
    private Double peso;
    private Boolean tieneChip;
    private Boolean estaLastimado;
    private String nombre;
    static private  Integer anioActual;


    public Perro(Boolean estaEnAdopcioon, String raza, Integer anioNacimiento, Double peso, Boolean tieneChip, Boolean estaLastimado, String nombre) {
        this.estaEnAdopcioon = estaEnAdopcioon;
        this.raza = raza;
        this.anioNacimiento = anioNacimiento;
        this.peso = peso;
        this.tieneChip = tieneChip;
        this.estaLastimado = estaLastimado;
        this.nombre = nombre;
    }

    public Perro(Boolean estaEnAdopcioon, String raza, Integer anioNacimiento, Double peso,  Boolean estaLastimado, String nombre) {
        this.estaEnAdopcioon = estaEnAdopcioon;
        this.raza = raza;
        this.anioNacimiento = anioNacimiento;
        this.peso = peso;
        this.estaLastimado = estaLastimado;
        this.nombre = nombre;
    }


//    Metodos
//    Edad
    public Integer edad(){
        return anioActual - anioNacimiento;
    }

//    puedePerderse
    public Boolean puedePerderse(){
        return !tieneChip;
    }

//    puedeAdoptarse
    public Boolean puedeAdoptarse(){
        return !estaLastimado && peso > 5.45;
    }

//    setAnioActual
    private static void setAnioActual(Integer anioActual) {
        Perro.anioActual = anioActual;
    }



//    Getters and Setters
    public Boolean getEstaEnAdopcioon() {
        return estaEnAdopcioon;
    }

    public void setEstaEnAdopcioon(Boolean estaEnAdopcioon) {
        this.estaEnAdopcioon = estaEnAdopcioon;
    }

    public String getRaza() {
        return raza;
    }

    public void setRaza(String raza) {
        this.raza = raza;
    }

    public Double getPeso() {
        return peso;
    }

    public void setPeso(Double peso) {
        this.peso = peso;
    }
}
