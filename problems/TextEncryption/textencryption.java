public class Main {
    public static void main(String[] args) throws Exception {
        Reader io = new Reader();
        Kattio out = new Kattio(System.in, System.out);

        while (true) {
            int n = io.nextInt();

            if (n == 0) {
                break;
            }

            String sentence = io.readLine();

            char[] charArr = sentence.replaceAll("\\W","").toUpperCase().toCharArray();
            char[] encoded = new char[charArr.length];

            int charIndex = 0;

            if (n > charArr.length) {
                String encodedWord = new String(charArr);
                out.print(encodedWord + "\n");
                out.flush();
            } else {
                for (int i = 0; i < n; i++) {
                    for (int j = i; j < charArr.length; j += n) {
                        encoded[j] = charArr[charIndex++];
                    }
                }

                String encodedWord = new String(encoded);
                out.print(encodedWord + "\n");
                out.flush();
            }
        }
    }
}
