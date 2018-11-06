import java.util.Iterator;
import java.util.LinkedList;

public class Main {

    //**********************************************************************************
    // THIS PROBLEM WAS DONE DURING A COMPETITION WHERE I HAD VERY LITTLE TIME TO
    // THEREFORE THE CODE-QUALITY IS LACKING
    //**********************************************************************************

    public static void main(String[] args) {

        Kattio io = new Kattio(System.in);
        int candies = io.getInt();

        int[] a = new int[candies];
        int[] b = new int[candies];
        boolean[][] diff = new boolean[candies][(2 * 100 * 100) + 1];

        for (int i = candies - 1; i > -1; i--) {
            a[i] = io.getInt();
        }

        for (int i = candies - 1; i > -1; i--) {
            b[i] = io.getInt();
        }

        if (candies == 1) {
            if (Math.abs(a[0]) <= Math.abs(b[0])) {
                System.out.println("A");
                System.exit(0);
            } else {
                System.out.println("B");
                System.exit(0);
            }
        }
        
        diff[0][10000 + a[0]] = true;
        diff[0][10000 - b[0]] = true;

        for (int i = 1; i < candies; i++) {
            for (int j = 0; j < 100*100+1; j++) {
                if (diff[i - 1][10000+j]) {
                    int alf = j+a[i];
                    int beata = j-b[i];

                    diff[i][10000 + alf] = true;
                    diff[i][10000 + beata] = true;
                }

                if (diff[i - 1][10000-j]) {
                    int alf = -j+a[i];
                    int beata = -j-b[i];


                    diff[i][10000 + alf] = true;
                    diff[i][10000 + beata] = true;
                }
            }
        }

        LinkedList<String> pos = new LinkedList<>();
        LinkedList<String> neg = new LinkedList<>();

        boolean found = false;

        for (int i = 0; i < 10000 + 1; i++) {
            if (diff[candies - 1][10000 + i]) {
                get_path(i, pos, candies, diff, a, b);
                found = true;
            }

            if (diff[candies - 1][10000 - i]) {
                get_path(-i, neg, candies, diff, a, b);
                found = true;
            }

            if (found) {
                break;
            }
        }


        if (pos.size() == neg.size()) {
            print(pos, neg);
        } else {
            if (neg.size() == 0) {
                print(pos);
            } else {
                print(neg);
            }
        }


    }

    static void get_path(int min, LinkedList<String> answer, int candies, boolean[][] diff, int[] a, int[] b) {
        for (int i = candies - 1; i > 0; i--) {
            if (diff[i - 1][10000 + min - a[i]]) {
                answer.addLast("A");
                min = min - a[i];
            } else if (diff[i - 1][10000 + min + b[i]]) {
                answer.addLast("B");
                min = min + b[i];
            }
        }

        if (min - a[0] == 0) {
            answer.addLast("A");
        } else if (min + b[0] == 0) {
            answer.addLast("B");
        }
    }

    static void print(LinkedList<String> a, LinkedList<String> b) {
        int dir = 0;

        Iterator<String> a_iter = a.iterator();
        Iterator<String> b_iter = b.iterator();

        while (a_iter.hasNext() && b_iter.hasNext()) {
            if (dir == -1) {
                System.out.print(b_iter.next());
            } else if (dir == 1) {
                System.out.print(a_iter.next());
            } else {
                String tmp_a = a_iter.next();
                String tmp_b = b_iter.next();

                if (tmp_a == tmp_b) {
                    System.out.print(tmp_b);
                } else if (tmp_a.equals("A")) {
                    System.out.print("A");
                    dir = 1;
                } else {
                    System.out.print("A");
                    dir = -1;
                }
            }
        }

    }

    static void print(LinkedList<String> answer) {
        for (String str : answer) {
            System.out.print(str);
        }
    }
}