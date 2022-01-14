package edu.zjnu;

import edu.zjnu.conf.Config;
import edu.zjnu.core.Server;
import edu.zjnu.core.HttpServlet;
import edu.zjnu.exception.ServerException;

/**
 * @description: 程序入口
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class BootStrap {

    public static void main(String[] args)  {
        try {
            Server server = new Server(Config.port ,new HttpServlet());
            server.start();
            System.out.println("服务器启动成功, 您现在可以访问 http://localhost:" + server.getPort());
        } catch (ServerException e) {
            System.out.println("服务器启动失败...");
            e.printStackTrace();
        }
//        System.in.read();
    }
}
