public class everywhere {
    public static void main(String[] args) {
        Kattio io = new Kattio(System.in);

        int n = io.getInt();
        String[][] cities = new String[n][];

        for (int i = 0; i < n; i++) {
            int size = io.getInt();
            cities[i] = new String[size];

            for (int j = 0; j < size; j++) {
                String temp = io.getWord();
                cities[i][j] = temp;
            }
        }

        for (int i = 0; i < cities.length; i++) {
            HashSet<String> answer = new HashSet<>();
            for (int j = 0; j < cities[i].length; j++) {
                answer.add(cities[i][j]);
            }

            io.println(answer.size());
            answer.clear();
        }
        io.close();
    }
}