/**
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

type CityInfoVo struct {
	models.City
	HaveChild bool `json:"haveChild"`
}
