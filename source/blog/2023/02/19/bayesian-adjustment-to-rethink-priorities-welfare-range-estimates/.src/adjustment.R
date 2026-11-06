## Packages

library(ggplot2)
library(ggthemes)

## Build the prior

min_95_ci = 1/5000
max_95_ci = 1/100

magic_constant = 1.6448536269514722
mean_lognormal = (log(min_95_ci) + log(max_95_ci))/2
sigma_lognormal = (log(max_95_ci) - log(min_95_ci))/ 
  ( 2 * magic_constant) 

num_points = 50000 ## 100000

xs = c(0:num_points)/num_points
ys = dlnorm(xs, meanlog = mean_lognormal, sdlog = sigma_lognormal)
ys = ys / sum(ys)
ys[1] = 2/3 ## so that it ends as 40% after normalization: 2/3 / (2/3 + 1) = 0.4
ys = ys / sum(ys)
cat("Checking that sum(ys) = 1")
cat(c("sum(ys) =", sum(ys)))

chicken_v_human = list()
chicken_v_human$xs = xs
chicken_v_human$ys = ys
chicken_v_human$color = c("red", rep("blue", num_points ))
chicken_v_human = as.data.frame(chicken_v_human)

## Plot the prior
title_text = "Prior over human vs chicken relative values"
subtitle_text=""
label_x_axis = "xs"
label_y_axis = "P(x)"

### Plot the prior over the whole x domain
ggplot(data=chicken_v_human, aes(x=xs, y=ys, color=color))+
  geom_point(size=0.5, aes(colour=color), show.legend = FALSE)+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=label_x_axis, 
    y=label_y_axis 
  ) +
  theme_tufte() + 
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5), 
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    plot.background=element_rect(fill = "white",colour = NA)
  )+
  scale_colour_manual(values = c("navyblue", "red"))

getwd() ## Working directory on which the file will be saved. Can be changed with setwd("/your/directory")
height = 5
width = floor(height*(1+sqrt(5))/2)
ggsave("prior.png", width=width, height=height)

### Plot the prior over only part of the x domain
subtitle_text="(zoomed in)"
ggplot(data=chicken_v_human, aes(x=xs, y=ys))+
  geom_point(color="navyblue", size=0.05)+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=label_x_axis, 
    y=label_y_axis 
  ) +
  theme_tufte() + 
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5), 
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    plot.background=element_rect(fill = "white",colour = NA)
  )+
  scale_x_continuous(limits = c(0, 0.05))+
  scale_y_continuous(limits = c(0, 0.005))

ggsave("prior-zoomed-in.png", width=width, height=height)

## Construct p(h|x)

p_w = 0.5
rp_estimate = 0.332

### Construct p(h|xW)

chicken_v_human$p_h_cond_x_W = rep(1/num_points, num_points+1)

### Construct p(h|x(not W))
are_within_one_order_of_magnitude = function(p1, p2){
  return(abs(log(p1/p2)/log(10)) <= 1)
}
are_within_one_order_of_magnitude(1,10)
are_within_one_order_of_magnitude(0.1,10)
are_within_one_order_of_magnitude(0.1,1)

start.time <- Sys.time()
count = 0
cache = TRUE
if(!cache){
  p_h_cond_x_not_W = c()
  for(x in (0:num_points)/num_points){
    print(count)
    count = count + 1
    is_within_one_oom_from_x = function(y){
      are_within_one_order_of_magnitude(x,y)
    }
    if(is_within_one_oom_from_x(rp_estimate)){
      num_close_to_x = sum(sapply(xs, is_within_one_oom_from_x))
      p_h_cond_x_not_W <- c(p_h_cond_x_not_W, 1/num_close_to_x)
    }else{
      p_h_cond_x_not_W <- c(p_h_cond_x_not_W, 0)
    }
  }
  end.time <- Sys.time()
  time.taken <- end.time - start.time
  time.taken
}

chicken_v_human$p_h_cond_x_not_W = p_h_cond_x_not_W

### Construct p(h|x) from p(h|xW) and p(h|x(not W))

chicken_v_human$p_h_cond_x = p_w * chicken_v_human$p_h_cond_x_W + (1-p_w)* chicken_v_human$p_h_cond_x_not_W

### plot p(h|x)
title_text = "P(h|x) update"
subtitle_text=""
label_x_axis = "xs"
label_y_axis = "P(h|x)"
ggplot(data=chicken_v_human, aes(x=xs, y=p_h_cond_x, color=color))+
  geom_point(size=0.05, aes(colour=color), show.legend = FALSE)+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=label_x_axis, 
    y=label_y_axis 
  ) +
  theme_tufte() + 
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5), 
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    plot.background=element_rect(fill = "white",colour = NA)
  )+
  scale_colour_manual(values = c("navyblue", "red"))

ggsave("p_h_x_update.png", width=width, height=height)

## Calculate p_h 

p_h = sum(chicken_v_human$ys * chicken_v_human$p_h_cond_x)
cat(p_h)

## Calculate p_x_given_h

chicken_v_human$p_x_given_h = chicken_v_human$ys * chicken_v_human$p_h_cond_x / p_h

## Plot p_x_given_h

title_text = "Posterior, P(x|h)"
subtitle_text=""
label_x_axis = "xs"
label_y_axis = "P(x|h)"
ggplot(data=chicken_v_human, aes(x=xs, y=p_x_given_h, color=color))+
  geom_point(size=0.5, aes(colour=color), show.legend = FALSE)+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=label_x_axis, 
    y=label_y_axis 
  ) +
  theme_tufte() + 
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5), 
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    plot.background=element_rect(fill = "white",colour = NA)
  )+
  scale_colour_manual(values = c("navyblue", "red"))

ggsave("posterior.png", width=width, height=height)

### Plot the posterior over only part of the x domain
subtitle_text="(zoomed in)"
ggplot(data=chicken_v_human, aes(x=xs, y=p_x_given_h))+
  geom_point(color="navyblue", size=0.05)+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=label_x_axis, 
    y=label_y_axis 
  ) +
  theme_tufte() + 
  theme(
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5), 
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=60, hjust=1),
    plot.background=element_rect(fill = "white",colour = NA)
  )+
  scale_x_continuous(limits = c(0, 0.1))+
  scale_y_continuous(limits = c(0, 0.005))

ggsave("posterior-zoomed-in.png", width=width, height=height)

## Show some indicators

chicken_v_human$p_x_given_h[1]

prior_average_value = sum(chicken_v_human$xs * chicken_v_human$ys)
posterior_average_value = sum(chicken_v_human$xs * chicken_v_human$p_x_given_h)
prior_average_value
posterior_average_value
