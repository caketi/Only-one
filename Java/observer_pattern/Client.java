public class Client {
    public static void main(String[] args) {
     Stock stock=new Stock(100.0);
     new Stocker("cake",stock);
     new Stocker("apple",stock);
     new Stocker("banana",stock);
     stock.setPrice(11.0);
     stock.setPrice(9.0);
     stock.setPrice(44.0);
 }
 }
 