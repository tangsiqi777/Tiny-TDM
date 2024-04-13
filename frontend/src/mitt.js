import Mitt from "mitt";

const emitter = Mitt();

//创建单例
class MittSingleInstance {
    //添加静态私有属性
    static #instanse;

    //添加静态方法
    static getInstanse() {
        //判断静态私有属性内是否有实例，如果没有就创建一个并返回，如果有就直接返回
        if (this.#instanse === undefined) {
            this.#instanse = new Mitt();
        }
        return this.#instanse;
    }
}

const tmitt = MittSingleInstance.getInstanse();

export  {tmitt}