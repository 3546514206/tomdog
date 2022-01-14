package edu.zjnu.core.http;

/**
 * @description: HttpServlet
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class HttpServlet {

    public void doGet(HttpRequest request, HttpResponse response) {
        System.out.println(request.getUrl());
        response.setBody("<html><h1>Hello World!</h1></html>");
    }

    public void doPost(HttpRequest request, HttpResponse response) {

    }
}
