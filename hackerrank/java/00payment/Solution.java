import java.util.*;
import java.text.*;

public class Solution {
    public static void main(String []args) {
        Scanner scanner = new Scanner(System.in);
        double payment = scanner.nextDouble();
        scanner.close();
        String[] nums = (Double.toString(payment)).split("\\.");
        int a = Integer.parseInt((nums[0]));
        System.out.println("US: $"  + a/1000 + "," + a%1000 + "." + nums[1].substring(0,2));
        System.out.println("India: Rs."  + a/1000 + "," + a%1000 + "." + nums[1]);
        System.out.println("China: ￥"  + a/1000 + "," + a%1000 + "." + nums[1]);
        System.out.println("France:"  + a/1000 + " " + a%1000 + "." + nums[1] + " €");
        nums[0].comp
        // System.out.println("India: Rs." + th + "," + payment);
        // System.out.println("China: " + th + "," + payment);
        // System.out.println("France: " + th + " " + payment + " ");
    }
}
