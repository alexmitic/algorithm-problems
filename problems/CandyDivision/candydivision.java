public class Main {

    public static void main(String[] args) {
        Kattio io = new Kattio(System.in);
        long amntCandy = io.getLong();
        TreeSet<Long> ans = new TreeSet<>();

        for (Long i = 2l; i <= Math.ceil(Math.sqrt((double)amntCandy)); i++) {
            if (amntCandy % i == 0) {
                ans.add(i);
                if (i != amntCandy / i) {
                    ans.add((amntCandy/i));
                }
            }
        }

        io.print(0);

        Iterator<Long> it = ans.iterator();
        while (it.hasNext()) {
            io.print(" " + (it.next() - 1));
        }

        io.print(" " + (amntCandy - 1));
        io.close();
    }
}