package edu.zjnu.core;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * @description: 请求映射
 * @author: 杨海波
 * @date: 2022-01-14
 **/
@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface RequestMapping {

    String url() default "";

    RequestMethod method() default RequestMethod.GET;
}
