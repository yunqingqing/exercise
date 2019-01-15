import org.junit.*;
import static org.junit.Assert.fail;

public class Singleton {
    private static volatile Singleton INSTANCE = null;

    private Singleton() {}

    public static Singleton getInstance() {
        if(INSTANCE == null) {
            synchronized(Singleton, class) {  //  第一次创建对象时才加锁
                if(INSTANCE == null) {
                    INSTANCE = new Singleton();
                }
            }
        }
        return INSTANCE;
    }
}
