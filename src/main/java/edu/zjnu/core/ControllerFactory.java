package edu.zjnu.core;

import edu.zjnu.biz.BaseController;
import edu.zjnu.biz.UserController;

import java.util.HashMap;
import java.util.Map;

/**
 * @description: 控制器
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class ControllerFactory {

    public static final Map<String, BaseController> controllerMap = new HashMap<>();

    static {
        UserController userController = new UserController();
        controllerMap.put("userController", userController);
    }
}
