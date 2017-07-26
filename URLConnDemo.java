import java.net.*;
import java.io.*;
import java.util.Scanner; 

public class URLConnDemo {

   public static void main(String [] args) {
     
     Scanner sc = new Scanner(System.in);
     System.out.println("Enter Service name : ");
     
     String input = sc.nextLine();
     String urlStr;
     String host = "http://127.0.0.1:8081/";
     
     if(input.equals("getservice2") || input.equals("getservice3"))
     {
       System.out.println("Enter Parameters : ");
       String para = sc.nextLine();
       para = para.replace(" ","%20");
       urlStr = host + input + "/" + para;
     }else{
       urlStr = host + input;
     }
     
     
      try {
         URL url = new URL(urlStr);
         URLConnection urlConnection = url.openConnection();
         HttpURLConnection connection = null;
         if(urlConnection instanceof HttpURLConnection) {
            connection = (HttpURLConnection) urlConnection;
         }else {
            System.out.println("Please enter an HTTP URL.");
            return;
         }
         
         BufferedReader in = new BufferedReader(
            new InputStreamReader(connection.getInputStream()));
         String urlString = "";
         String current;
         
         while((current = in.readLine()) != null) {
            urlString += current;
         }
         System.out.println(urlString);
      }catch(IOException e) {
         e.printStackTrace();
      }
   }
}