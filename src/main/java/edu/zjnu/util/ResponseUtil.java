package edu.zjnu.util;

import com.alibaba.fastjson.JSON;

import java.util.List;

/**
 * @description: ResponseUtil
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class ResponseUtil {

    public static String hand(Object o) {
        if (o instanceof List) {
            return JSON.toJSONString(o);
        }

        return null;
    }
}
