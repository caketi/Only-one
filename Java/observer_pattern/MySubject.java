import java.util.*;
public abstract class MySubject {
  private Vector<MyObserver>vector = new Vector<MyObserver>();
  public void addMyOberserver(MyObserver Observer){
	  this.vector.add(Observer);
  } 
  public void removeMyOberserver(MyObserver Observer){
	  this.vector.remove(Observer);
  } 
  public void Notice(){
	  for (int i=0;i<vector.size();i++){
		  ((MyObserver)vector.elementAt(i)).response();
	  }
  }
}
