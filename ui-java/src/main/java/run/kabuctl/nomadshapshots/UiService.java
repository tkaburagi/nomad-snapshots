package run.kabuctl.nomadshapshots;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.cloud.client.loadbalancer.LoadBalancerClient;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
public class UiService {

    private final RestTemplate restTemplate;
    private final ObjectMapper objectMapper;
    LoadBalancerClient loadBalancerClient;


    public UiService(RestTemplateBuilder builder, ObjectMapper objectMapper, LoadBalancerClient loadBalancerClient) {
        this.restTemplate = builder.build();
        this.objectMapper = objectMapper;
        this.loadBalancerClient = loadBalancerClient;
    }


    public User[] getAllUsers() throws Exception {
        String url = getUrl();
        String result = this.restTemplate.getForObject(url + "/users", String.class);
        System.out.println(result);
        User[] userList = this.objectMapper.readValue(result, User[].class);

        return userList;
    }

    public String getVersion() throws Exception {
        String url = getUrl();
        String result = this.restTemplate.getForObject(url + "/version", String.class);
        System.out.println(result);
        Version version = this.objectMapper.readValue(result, Version.class);

        return version.getVersion();
    }


    public String getUrl() {
        return this.loadBalancerClient.choose("api-go").getUri().toString();
    }


}
