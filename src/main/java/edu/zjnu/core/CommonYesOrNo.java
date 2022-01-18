package edu.zjnu.core;

/**
 * @description: 是否，有无，在或不在
 * @author: 杨海波
 * @date: 2022-01-18
 **/
public enum CommonYesOrNo {

    YES(1, "是"), NO(0, "否");

    private Integer code;
    private String desc;

    CommonYesOrNo(Integer code, String desc) {
        this.code = code;
        this.desc = desc;
    }


    public Integer getCode() {
        return code;
    }

    public String getDesc() {
        return desc;
    }

    /**
     * 根据code获取去value
     * @param code
     * @return
     */
    public static CommonYesOrNo getEnumByCode(Integer code) {
        for (CommonYesOrNo c : CommonYesOrNo.values()) {
            if (code.equals(c.getCode())) {
                return c;
            }
        }
        return null;
    }
}
