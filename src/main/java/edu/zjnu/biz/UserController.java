package edu.zjnu.biz;

import edu.zjnu.core.RequestMethod;
import edu.zjnu.core.RequestMapping;

import java.util.List;
import java.util.Map;

/**
 * @description: UserController
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class UserController extends BaseController {

    @RequestMapping(url = "/getUserList", method = RequestMethod.POST)
    public List<User> getUserList(Map param) {
        List<User> userList = (List<User>) DataBase.DATABASE.get("userList");
        userList.forEach(e->{
            System.out.println(e.toString());
        });
        return userList;
    }
}
