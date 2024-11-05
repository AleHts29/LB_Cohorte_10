public class main {
    public static void main(String[] args) {

        Perro newPerro = new Perro(
                false,
                "perro",
                2019,
                40.0,
//                false,
                false,
                "Happy"
        );



        System.out.println(newPerro.puedeAdoptarse());

    }
}
