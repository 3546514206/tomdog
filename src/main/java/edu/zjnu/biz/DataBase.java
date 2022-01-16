package edu.zjnu.biz;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @description: 数据库
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class DataBase {

    public static final Map<String, Object> DATABASE = new HashMap<>();

    static {
        User sheng = new User(1L, "shengxinyi");
        User zhao = new User(2L, "zhaomeiyu");
        User li = new User(3L, "lisongpeng");
        User yang = new User(4L, "yanghaibo");
        List<User> userList = new ArrayList<>();
        userList.add(sheng);
        userList.add(zhao);
        userList.add(li);
        userList.add(yang);
        DATABASE.put("userList", userList);
    }


}
