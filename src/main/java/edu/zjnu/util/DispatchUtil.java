package edu.zjnu.util;

import edu.zjnu.biz.BaseController;
import edu.zjnu.core.*;
import edu.zjnu.exception.ServerException;

import java.lang.annotation.Annotation;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.HashMap;
import java.util.Map;

/**
 * @description: 转发工具
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class DispatchUtil {

    public static String dispatchUtil(HttpRequest request, HttpResponse response) throws ServerException {
        String url = request.getUrl();
        String requestMethod = request.getMethod();

        if (null == url) return "";

        Map<String, BaseController> controllerMap = ControllerFactory.controllerMap;

        for (BaseController controller : controllerMap.values()) {
            // 获取所有的业务接口
            java.lang.reflect.Method[] methods = controller.getClass().getMethods();
            //根据注解获取到请求的接口并执行
            for (Method method : methods) {
                RequestMapping requestMapping = method.getAnnotation(RequestMapping.class);
                System.out.println("即将请求业务接口");

                if (requestMapping.url().equalsIgnoreCase(url)
                        && requestMapping.method().getRequestMethod().equalsIgnoreCase(requestMethod)) {
                    try {
                        Object o = method.invoke(controller, new HashMap<>());
                        //处理返回的业务数据
                        return ResponseUtil.hand(o);
                    } catch (IllegalAccessException | InvocationTargetException e) {
                        e.printStackTrace();
                        throw new ServerException("方法转发失败");
                    }
                }
                System.out.println("业务接口处理完成");

            }

        }

        return null;
    }
}
