public class Stock extends MySubject{
    private double price;
    public Stock(double price){
        this.price=price;
    }   
    public double getPrice(){
        return price;
    }
    public void setPrice(double newPrice){	 
        System.out.println("************************");
        if((newPrice-this.price)/this.price>0.05||(newPrice-this.price)/this.price<-0.05){
            //this.price=price;
           // System.out.println("************************");
            System.out.println("股票的价格变化超过5%，变化率为"+(newPrice-this.price)/this.price*100+"%"+"\n现在的价格为"+newPrice);
            super.Notice();
        }else{
            System.out.println("股票的价格变化没超过5%，变化率为"+(newPrice-this.price)/this.price*100+"%"+"\n现在的价格为"+newPrice);
        }	   
    }
} 
