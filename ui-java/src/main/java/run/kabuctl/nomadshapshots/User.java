package run.kabuctl.nomadshapshots;

import java.io.Serializable;

public class User implements Serializable {

    String id;
    String username;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

}
