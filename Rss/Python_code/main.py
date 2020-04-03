import pymysql
import datetime


def initDB():
    db = pymysql.connect("localhost", "root", "123456", "movierss")
    return db

def initMovie_dat():
    db = initDB()
    cursor = db.cursor()
    file = open("./data/ml-1m/movies.dat","rb")
    while 1:
        lines = file.readlines(100000)
        if not lines:
            break
        for line in lines:
            item = line.decode("unicode_escape")
            movie_id_index = item.find('::')
            movie_id = item[:movie_id_index]
            movie_title_index = item.rfind('::')
            movie_title = item[movie_id_index+2:movie_title_index]
            movie_genres = item[movie_title_index+2:]
            sql = "INSERT INTO movie(movie_id,title,genres) VALUES  ('%d','%s','%s')" % (int(movie_id),movie_title,movie_genres)
            try:
                # 执行SQL语句
                cursor.execute(sql)
                # 向数据库提交
                db.commit()
            except:
                # 发生错误时回滚
                db.rollback()
    db.close()
    file.close()

def initRatings_dat():
    db = initDB()
    cursor = db.cursor()
    file = open("./data/ml-1m/ratings.dat","rb")
    while 1:
        lines = file.readlines(100000)
        if not lines:
            break
        for line in lines:
            item = line.decode("unicode_escape")
            user_id_index = item.find('::')
            user_id = item[:user_id_index]
            movie_id_index = item.find('::',user_id_index+2)
            movie_id = item[user_id_index+2:movie_id_index]
            rating_index = item.find('::',movie_id_index+2)
            rating = item[movie_id_index+2:rating_index]
            timestamp = item[rating_index+2:]
            dateArray = datetime.datetime.utcfromtimestamp(float(timestamp))
            otherStyleTime = dateArray.strftime("%Y-%m-%d %H:%M:%S")
            sql = "INSERT INTO ratings(user_id,movie_id,rating,created_time) VALUES  ('%d','%d','%d','%s')" % (int(user_id),int(movie_id),int(rating),otherStyleTime)
            try:
                # 执行SQL语句
                cursor.execute(sql)
                # 向数据库提交
                db.commit()
            except:
                # 发生错误时回滚
                db.rollback()
    db.close()
    file.close()

def initUsers_dat():
    db = initDB()
    cursor = db.cursor()
    file = open("./data/ml-1m/users.dat","rb")
    while 1:
        lines = file.readlines(100000)
        if not lines:
            break
        for line in lines:
            item = line.decode("unicode_escape")
            user_id_index = item.find('::')
            user_id = item[:user_id_index]
            gender_index = item.find('::',user_id_index+2)
            gender = item[user_id_index+2:gender_index]
            age_index = item.find('::',gender_index+2)
            age = item[gender_index+2:age_index]
            occoupation_index = item.find('::', age_index + 2)
            occoupation = item[age_index + 2:occoupation_index]
            zip_code = item[occoupation_index + 2:]
            sql = "INSERT INTO rated_users(user_id,gender,age,occupation,zip_code) VALUES  ('%d','%s','%d','%s','%s')" % (int(user_id),gender,int(age),occoupation,zip_code)
            try:
                # 执行SQL语句
                cursor.execute(sql)
                # 向数据库提交
                db.commit()
            except:
                # 发生错误时回滚
                db.rollback()
    db.close()
    file.close()

def main():
    print("init DB beginning")
    initMovie_dat()
    initRatings_dat()
    initUsers_dat()
    print("init DB end")


if __name__ == '__main__':
    main()