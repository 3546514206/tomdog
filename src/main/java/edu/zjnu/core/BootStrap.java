package edu.zjnu.core;

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

    public void load() {
        try {
            System.out.println("服务器配置： 端口[" + Config.port + "]");
            Server server = new Server(Config.port, new HttpServlet());
            server.start();
            System.out.println("服务器启动成功");
            Class.forName("edu.zjnu.biz.DataBase");
            Class.forName("edu.zjnu.core.ControllerFactory");
        } catch (ServerException | ClassNotFoundException e) {
            System.out.println("服务器启动失败...");
            e.printStackTrace();
        }
//        System.in.read();
    }
}
