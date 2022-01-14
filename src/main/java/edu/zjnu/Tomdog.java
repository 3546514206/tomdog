package edu.zjnu;

import edu.zjnu.core.BootStrap;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Tomdog {

	public static void main(String[] args) {
		new BootStrap().load();
		SpringApplication.run(Tomdog.class, args);
	}

}
