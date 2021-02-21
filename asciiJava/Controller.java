package com.example.ascii;

import com.example.ascii.Logic.Handler;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;

@RestController
@RequestMapping("/ascii")
public class Controller {
    Handler handler = new Handler();
    @GetMapping("/render/{filename}/{width}/{inverse}")
    public void render(@PathVariable String filename,@PathVariable int width,@PathVariable boolean inverse) throws IOException {
        String readFrom = "E:\\Go\\asciiServer\\images\\"+filename+".png";
        String saveTo = "E:\\Go\\asciiServer\\textfiles\\"+filename+".txt";
        handler.handle(width,readFrom,saveTo,inverse);
    }
    @GetMapping("/render/{filename}/{width}")
    public void render(@PathVariable String filename,@PathVariable int width) throws IOException {
        String readFrom = "E:\\Go\\asciiServer\\images\\"+filename+".png";
        String saveTo = "E:\\Go\\asciiServer\\textfiles\\"+filename+".txt";
        handler.handle(width,readFrom,saveTo);
    }
}
