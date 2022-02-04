package edu.zjnu.biz;

import edu.zjnu.core.RequestMethod;
import edu.zjnu.core.RequestMapping;
import edu.zjnu.exception.ServerException;

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
        userList.forEach(e -> {
            System.out.println(e.toString());
        });
        return userList;
    }

    @RequestMapping(url = "/getUserById", method = RequestMethod.POST)
    public User getUserById(Map param) throws ServerException {

        List<User> users = (List) DataBase.DATABASE.values();
        Long userId = (Long) param.get("userId");

        if (null == userId) {
            throw new ServerException("参数为空");
        }

        User rs = null;
        for (User user : users) {
            if (userId.equals(user.getUserId())) {
                rs = user;
                break;
            }
        }

        return rs;
    }
}


