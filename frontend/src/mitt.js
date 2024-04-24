import Mitt from "mitt";

//创建单例
class MittSingleInstance {
    //添加静态私有属性
    static mittInstance;

    //添加静态方法
    static getInstance() {
        //判断静态私有属性内是否有实例，如果没有就创建一个并返回，如果有就直接返回
        if (this.mittInstance === undefined) {
            this.mittInstance = new Mitt();
        }
        return this.mittInstance;
    }
}

const SingleMitt = MittSingleInstance.getInstance();

export {SingleMitt}