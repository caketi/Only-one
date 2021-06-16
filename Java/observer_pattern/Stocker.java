public class Stocker implements MyObserver{
    private String name;
    public Stock  stock;	
	public Stocker(String name,Stock Mysubject){
		this.name=name;
		Mysubject.addMyOberserver(this);
	}
	@Override
	public void response() {
		
		System.out.println(name+"被通知股票发生了变化");
	}	
}

