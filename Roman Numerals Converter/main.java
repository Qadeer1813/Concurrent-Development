// Stduent Name: Qadeer Hussain
// Student ID: C00270632
// Module: Concurrent Development
import java.util.Scanner;

public class main{
    public static int getRomanChar(char romanChar) {
        romanChar = Character.toUpperCase(romanChar);
        
        if (romanChar == 'I') {
            return 1;
        } else if (romanChar == 'V') {
            return 5;
        } else if (romanChar == 'X') {
            return 10;
        } else if (romanChar == 'L') {
            return 50;
        } else if (romanChar == 'C') {
            return 100;
        } else if (romanChar == 'D') {
            return 500;
        } else if (romanChar == 'M') {
            return 1000;
        }
        return romanChar;
    }

    public static int romanToInt(String s) {

        int total = 0;  // The total is what the roman numeral is as a intger
        int previousValue = 0; // This tracks the previous roman numeral 
        
        // Go to through character from right to left
        for (int i = s.length() - 1; i >= 0; i--) {
            // Convert the current roman numeral to its integer value
            int currentValue = getRomanChar(s.charAt(i));
            // If the current roman numeral as a integer is less than the previous roman numeral it will subtract is example: IV = 5 - 1
            if (currentValue < previousValue) {
                total = total - currentValue;
            } 
            else {
            //  Otherwsie just add it to the total
                total = total + currentValue;
            }
            
            // Update the the previous value to the current value 
            previousValue = currentValue;
        }
        
        // Return the total after converting
        return total;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Input: ");
        String input = scanner.nextLine();
        // Convert the roman numeral string to a integer
        int result = romanToInt(input);
        System.out.println("Output: " + result);
    }     
}