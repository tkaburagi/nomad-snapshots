package run.kabuctl.nomadshapshots;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class UiController {

    private final UiService uiService;

    public UiController(UiService uiService) {
        this.uiService = uiService;
    }

    @GetMapping(value = "/")
    public String home(Model model) throws Exception {

        model.addAttribute("users", this.uiService.getAllUsers());
        return "ui/index";
    }

    @GetMapping(value = "/api-location")
    public String getApiLocation(Model model) throws Exception {

        model.addAttribute("api", this.uiService.getUrl());
        model.addAttribute("version", this.uiService.getVersion());
        return "ui/demo";
    }

}

