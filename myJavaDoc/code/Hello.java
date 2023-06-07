import java.util.Scanner;

public class Hello {
	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		System.out.println("Insert one digit number");
		String input = scanner.nextLine();
		int num = Integer.parseInt(input);
		System.out.println("received value: " + input);
		System.out.printf("num=%d%n", num);
		scanner.close();
	}
}
