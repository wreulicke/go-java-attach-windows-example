

class test {
  public static void main(String[] args) throws Throwable {
    new Thread(() -> {
      while(true) {
        try {
          Thread.sleep(300);
          System.out.println("test");
        } catch(Exception e) {
        }
      }
    }).start();
  }
} 
