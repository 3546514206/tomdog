package edu.zjnu.core;

import edu.zjnu.exception.ServerException;
import edu.zjnu.util.DispatchUtil;

import java.util.HashMap;
import java.util.Map;

/**
 * @description: HttpServlet
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class HttpServlet {

    public void doGet(HttpRequest request, HttpResponse response) {
        System.out.println(request.getUrl());

    }

    public void doPost(HttpRequest request, HttpResponse response) {
        String body = null;
        try {
            body = DispatchUtil.dispatchUtil(request,response);
        } catch (ServerException e) {
            e.printStackTrace();
        }

        response.setBody(body);
    }
}
