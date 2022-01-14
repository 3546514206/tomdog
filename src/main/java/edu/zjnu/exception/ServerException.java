package edu.zjnu.exception;

/**
 * @description: todo
 * @author: 杨海波
 * @date: 2022-01-14
 **/
public class ServerException extends Exception {


    public ServerException(String message) {
        super(message);
    }

    @Override
    public String toString() {
        return "ServerException{" +
                "message='" + getMessage() + '\'' +
                '}';
    }
}
